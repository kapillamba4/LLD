package usecases

import (
	"encoding/csv"
	"ethledger/app/models"
	"os"
)

type CSVWriter struct{}

func (w *CSVWriter) WriteToFile(txns []models.EnrichedTransaction, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{
		"Transaction Hash",
		"Date & Time",
		"From Address",
		"To Address",
		"Transaction Type",
		"Asset Contract Address",
		"Asset Symbol / Name",
		"Token ID",
		"Value / Amount",
		"Gas Fee (ETH)",
		"Category",
	}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, tx := range txns {
		row := []string{
			tx.Hash,
			tx.Timestamp.Format("2006-01-02 15:04:05"),
			tx.From,
			tx.To,
			tx.Type,
			tx.AssetContract,
			tx.AssetSymbol,
			tx.TokenID,
			tx.Amount,
			tx.GasFeeEth,
			tx.Category,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
