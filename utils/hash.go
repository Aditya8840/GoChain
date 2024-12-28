package utils

import (
	"crypto/sha256"
	"strconv"

	"github.com/Aditya8840/GoChain/types"
)


func CalculateHash(block types.Block) []byte {
	record := strconv.Itoa(block.Index)+block.Timestamp+string(block.Data[:])+string(block.Prev_hash[:])
	hash := sha256.Sum256([]byte(record))
	return hash[:]
}