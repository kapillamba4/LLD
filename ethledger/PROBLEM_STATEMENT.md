# Low Level Design - Ethereum Transaction Exporter

## Problem

Build a tool that fetches Ethereum wallet transactions and exports them to CSV.

### Requirements

**Input**: Ethereum wallet address  
**Output**: CSV file with transaction history  

### Transaction Types to Support
- External transfers (wallet-to-wallet)
- Internal transfers (smart contract calls)
- ERC-20 token transfers
- ERC-721 NFT transfers

### CSV Fields Required
```
Transaction Hash, Date, From, To, Type, Contract Address, 
Symbol, Token ID, Amount, Gas Fee (ETH)
```

### Technical Constraints
- Use Etherscan/Alchemy/Blockscout API (2 req per second rate limit)
- Handle pagination
- Handle all type of transactions
