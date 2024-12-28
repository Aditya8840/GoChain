package types

import (
	"time"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Index int
	Timestamp string
	Hash []byte
	Data []byte
	Prev_hash []byte
}

func CalculateHash(block Block) []byte {
	record := strconv.Itoa(block.Index)+block.Timestamp+string(block.Data[:])+string(block.Prev_hash[:])
	hash := sha256.Sum256([]byte(record))
	return hash[:]
}

func CreateBlock(old *Block, data string) *Block {
	var block Block
	currentTime := time.Now()
	block.Index = old.Index + 1
	block.Timestamp = currentTime.String()
	block.Data = []byte(data)
	block.Prev_hash = old.Hash
	block.Hash = CalculateHash(block)
	return &block
}

func Genesis() *Block {
	block := &Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Hash:      nil,
		Data:      []byte("Genesis Block"),
		Prev_hash: nil,
	}
	block.Hash = CalculateHash(*block)
	return block
}