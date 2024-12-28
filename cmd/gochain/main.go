package main

import (
	"fmt"

	"github.com/Aditya8840/GoChain/types"
)


func main() {
	chain := &types.Chain{}
	chain.InitChain()
	fmt.Println(chain)

	chain.AddBlock("1st block")
	fmt.Println(chain)

	fmt.Println(chain.IsValid())
}