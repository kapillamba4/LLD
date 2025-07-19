package models

import "time"

type TransactionsProviderResponse struct {
	External []Transaction
	Internal []Transaction
	ERC20    []Transaction
	ERC721   []Transaction
}

type Transaction struct {
	Hash          string    `csv:"Transaction Hash"`
	Timestamp     time.Time `csv:"Date & Time"`
	From          string    `csv:"From Address"`
	To            string    `csv:"To Address"`
	Type          string    `csv:"Transaction Type"`
	AssetContract string    `csv:"Asset Contract Address"`
	AssetSymbol   string    `csv:"Asset Symbol / Name"`
	TokenID       string    `csv:"Token ID"`
	Amount        string    `csv:"Value / Amount"`
	GasFeeEth     string    `csv:"Gas Fee (ETH)"`
}
