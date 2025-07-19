package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"ethledger/app/models"
	"ethledger/app/usecases"

	"golang.org/x/time/rate"
)

var _ usecases.TransactionsProvider = (*EtherscanClient)(nil)

type EtherscanClient struct {
	APIKey  string
	limiter *rate.Limiter
}

func NewEtherscanClient(apiKey string) *EtherscanClient {
	limiter := rate.NewLimiter(rate.Limit(1), 1)
	return &EtherscanClient{
		APIKey:  apiKey,
		limiter: limiter,
	}
}

func (e *EtherscanClient) Fetch(walletAddress string) (resp *models.TransactionsProviderResponse, err error) {
	resp = &models.TransactionsProviderResponse{}
	resp.External, err = e.fetchTxListForAllPages(walletAddress, "txlist")
	if err != nil {
		return nil, fmt.Errorf("fetch external txs: %w", err)
	}

	resp.Internal, err = e.fetchTxListForAllPages(walletAddress, "txlistinternal")
	if err != nil {
		return nil, fmt.Errorf("fetch internal txs: %w", err)
	}

	resp.ERC20, err = e.fetchTxListForAllPages(walletAddress, "tokentx")
	if err != nil {
		return nil, fmt.Errorf("fetch erc20 txs: %w", err)
	}

	resp.ERC721, err = e.fetchTxListForAllPages(walletAddress, "tokennfttx")
	if err != nil {
		return nil, fmt.Errorf("fetch erc721 txs: %w", err)
	}

	return resp, nil
}

func (e *EtherscanClient) fetchTxListForAllPages(wallet string, action string) (result []models.Transaction, err error) {
	startBlock := 0
	endBlock := 99999999

	for {
		etherscanResp, transactions, err := e.fetchTxListByBlocks(wallet, action, startBlock, endBlock)
		if err != nil {
			return nil, err
		}

		if len(transactions) == 0 {
			break
		}

		for _, t := range etherscanResp.Result {
			blockNumber, err := strconv.ParseInt(t.BlockNumber, 10, 64)
			if err != nil {
				return nil, err
			}
			startBlock = max(int(blockNumber+1), startBlock)
		}
		result = append(result, transactions...)
	}

	return result, nil
}

func (e *EtherscanClient) fetchTxListByBlocks(wallet string, action string, startBlock int, endBlock int) (etherscanResp *EtherscanTxListResponse, txs []models.Transaction, err error) {
	err = e.getParsedResponseByBlocks(action, wallet, startBlock, endBlock, &etherscanResp)
	if err != nil {
		return nil, nil, err
	}
	if etherscanResp.Status != "1" {
		if etherscanResp.Message == "No transactions found" {
			return nil, nil, nil
		}
		return nil, nil, fmt.Errorf("etherscan API error: %s", etherscanResp.Message)
	}

	for _, r := range etherscanResp.Result {
		ts, _ := strconv.ParseInt(r.TimeStamp, 10, 64)
		gasUsed, _ := strconv.ParseFloat(r.GasUsed, 64)
		gasPriceWei, _ := strconv.ParseFloat(r.GasPrice, 64)
		valueWei, _ := strconv.ParseFloat(r.Value, 64)
		gasFeeEth := gasUsed * gasPriceWei / 1e18 // 1 ETH = 1e18 Wei
		amountEth := valueWei / 1e18              // Transaction Amount in ETH

		txs = append(txs, models.Transaction{
			Hash:          r.Hash,
			Timestamp:     time.Unix(ts, 0),
			From:          r.From,
			To:            r.To,
			Type:          "ETH",
			AssetContract: r.ContractAddress,
			AssetSymbol:   r.TokenSymbol,
			TokenID:       r.TokenID,
			Amount:        fmt.Sprintf("%.8f", amountEth),
			GasFeeEth:     fmt.Sprintf("%.8f", gasFeeEth),
		})
	}
	return etherscanResp, txs, nil
}

func (e *EtherscanClient) getParsedResponseByBlocks(action string, wallet string, startBlock, endBlock int, target interface{}) error {
	ctx := context.Background()
	err := e.limiter.Wait(ctx)
	if err != nil {
		return fmt.Errorf("rate limiter error: %w", err)
	}

	url := fmt.Sprintf("https://api.etherscan.io/api?module=account&action=%s&address=%s&page=1&offset=10000&startblock=%d&endblock=%d&sort=asc&apikey=%s", action, wallet, startBlock, endBlock, e.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, target); err != nil {
		fmt.Println(string(body))
		return err
	}

	return nil
}

type EtherscanTxListResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Result  []EtherscanTxRecord `json:"result"`
}

type EtherscanTxRecord struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenID           string `json:"tokenID"`
}
