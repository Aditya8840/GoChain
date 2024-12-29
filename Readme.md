# GoChain

A simple blockchain implementation in Go that demonstrates the basic concepts of blockchain technology using SHA-256 hashing.

## Features

- Basic block structure with index, timestamp, data, and hash values
- SHA-256 hashing algorithm implementation
- Block validation mechanism
- Chain integrity verification

## Block Structure

```go
type Block struct {
    Index     int
    Timestamp string
    Hash      []byte
    Data      []byte
    Prev_hash []byte
}
```

## Core Functions

- **CalculateHash**: Generates SHA-256 hash for a block using its properties
- **CreateBlock**: Creates a new block with given data and previous block's hash
- **IsBlockValid**: Validates a block by checking its index, previous hash, and current hash

## Security Feature

The implementation includes a validation mechanism that ensures chain integrity. If any data in a block is modified, it will invalidate the entire chain from that point forward due to hash mismatches.

**This is a basic implementation of blockchain that will be enhanced with more features and security measures.**