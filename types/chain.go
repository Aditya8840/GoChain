package types

import (
	"bytes"

	"github.com/Aditya8840/GoChain/utils"
)

type Chain struct {
	Blocks []*Block
}


func (c *Chain) AddBlock(data string) {
    var newBlock = CreateBlock(c.Blocks[len(c.Blocks)-1], data)
    c.Blocks = append(c.Blocks, newBlock)
}

func (c *Chain) InitChain(){
	c.Blocks = append(c.Blocks, Genesis())
}

func (c *Chain) IsValid() bool {
	for i:=1; i<len(c.Blocks); i++ {
		currentBlock := c.Blocks[i]
        prevBlock := c.Blocks[i-1]

		if currentBlock.Index - 1 != prevBlock.Index {
			return false
		}

        if !bytes.Equal(currentBlock.Prev_hash, prevBlock.Hash) {
			return false
		}

		if !bytes.Equal(currentBlock.Hash, utils.CalculateHash(currentBlock)) {
            return false
        }
	}
	return true
}