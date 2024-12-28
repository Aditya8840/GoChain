package types

import (
	"time"

	"github.com/Aditya8840/GoChain/utils"
)

type Block struct {
	Index int
	Timestamp string
	Hash []byte
	Data []byte
	Prev_hash []byte
}

func CreateBlock(old *Block, data string) *Block {
	var block Block
	currentTime := time.Now()
	block.Index = old.Index + 1
	block.Timestamp = currentTime.String()
	block.Data = []byte(data)
	block.Prev_hash = old.Hash
	block.Hash = utils.CalculateHash(block)
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
	block.Hash = utils.CalculateHash(block)
	return block
}