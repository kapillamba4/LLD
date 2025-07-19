package usecases

import "ethledger/app/models"

type Transformer struct{}

const (
	CategoryExternalTransfer = "external"
	CategoryInternalTransfer = "internal"
	CategoryERC20Transfer    = "erc20"
	CategoryERC721Transfer   = "erc721"
)

func (t *Transformer) Transform(txns []models.Transaction, category string) []models.EnrichedTransaction {
	var enriched []models.EnrichedTransaction

	for _, tx := range txns {
		enriched = append(enriched, models.EnrichedTransaction{
			Transaction: tx,
			Category:    category,
		})
	}

	return enriched
}
