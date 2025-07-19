package usecases

import "ethledger/app/models"

type TransactionsProvider interface {
	Fetch(walletAddress string) (*models.TransactionsProviderResponse, error)
}
