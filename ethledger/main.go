package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"ethledger/app/services"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}

	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	wallet := exportCmd.String("wallet", "", "Ethereum wallet address")
	outfile := exportCmd.String("outfile", "txns.csv", "Output CSV file")
	if len(os.Args) < 2 {
		fmt.Println("expected 'export' command")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "export":
		exportCmd.Parse(os.Args[2:])
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	ledgerService := services.NewLedgerService(apiKey)
	if exportCmd.Parsed() {
		if *wallet == "" {
			log.Fatal("--wallet address is required")
		}

		fmt.Println("Fetching transactions for wallet:", *wallet)
		err := ledgerService.ExportToCSV(*wallet, *outfile)
		if err != nil {
			log.Fatal("Error exporting to CSV:", err)
		}
		fmt.Println("Successfully exported transactions to:", *outfile)
	}
}
