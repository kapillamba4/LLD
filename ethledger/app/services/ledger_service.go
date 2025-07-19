package services

import (
	"ethledger/app/models"
	"ethledger/app/providers"
	"ethledger/app/usecases"
)

type LedgerService struct {
	txnProvider usecases.TransactionsProvider
	transformer *usecases.Transformer
	csvWriter   *usecases.CSVWriter
}

func NewLedgerService(apiKey string) *LedgerService {
	return &LedgerService{
		txnProvider: providers.NewEtherscanClient(apiKey),
		transformer: &usecases.Transformer{},
		csvWriter:   &usecases.CSVWriter{},
	}
}

func (l *LedgerService) getAllTransactions(walletAddr string) ([]models.EnrichedTransaction, error) {
	resp, err := l.txnProvider.Fetch(walletAddr)
	if err != nil {
		return nil, err
	}

	enrichedTransactions := []models.EnrichedTransaction{}

	// External(Normal) Transfers - These are direct transfers between user controlled addresses.
	enrichedTransactions = append(enrichedTransactions, l.transformer.Transform(resp.External, usecases.CategoryExternalTransfer)...)
	// Internal Transfers - These are transfers that occur within smart contracts & not directly initiated by users.
	enrichedTransactions = append(enrichedTransactions, l.transformer.Transform(resp.Internal, usecases.CategoryInternalTransfer)...)
	// Token Transfers
	enrichedTransactions = append(enrichedTransactions, l.transformer.Transform(resp.ERC20, usecases.CategoryERC20Transfer)...)
	enrichedTransactions = append(enrichedTransactions, l.transformer.Transform(resp.ERC721, usecases.CategoryERC721Transfer)...)

	return enrichedTransactions, nil
}

func (l *LedgerService) ExportToCSV(walletAddr string, outPath string) error {
	enrichedTransactions, err := l.getAllTransactions(walletAddr)
	if err != nil {
		return err
	}

	err = l.csvWriter.WriteToFile(enrichedTransactions, outPath)
	if err != nil {
		return err
	}

	return nil
}
