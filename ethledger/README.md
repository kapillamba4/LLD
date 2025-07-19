# EthLedger - Transaction Exporter

A Go-based command-line tool for exporting Ethereum wallet transactions to CSV format.

## Quick Start

### Prerequisites

- Go 1.23.0 or higher
- Etherscan API key (get one at [etherscan.io/apis](https://etherscan.io/apis))

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd ethledger
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create a `.env` file with your Etherscan API key:
```bash
echo "API_KEY=your_api_key_here" > .env
```

4. Build the application:
```bash
go build -o ethledger main.go
```

5. Test the application end to end.
```bash
go test -v ./tests/e2e/ -run TestEthledgerErrorHandling
```
### Usage

Export transactions for a wallet address:
```bash
./ethledger export --wallet <wallet_address> --outfile txns.csv
```

## Project Structure

```
ethledger/
├── main.go                          # Entry point
├── go.mod                           # Go module definition
├── .env                             # Environment variables
├── README.md                        # This file
└── app/
    ├── models/                      # Data structures
    ├── providers/                   # External data sources
    ├── services/                    # Application Business logic
    └── usecases/                    # Domain specific use cases
```

## Architecture Decisions

### 1. Used TransactionsProvider interface
- **Rationale**: Enables easy testing, mocking, and swapping of implementations (Etherscan, Alchemy , Blockscout, Infura, etc)

### 2. Implement client-side request throttling 
- **Rationale**: Free plan has 2 req/sec limit

### 3. Separation of Data Fetching and Processing
- **Rationale**: Single Responsibility Principle, easier testing and future extensibility
