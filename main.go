package main

import (
	"fmt"
	"github.com/sarthak815/Blockchain-Capstone/core"
)

func main() {
	fmt.Println("Creating Genesis Block")
	chain := core.InitBlockChain()
	fmt.Println("Creating Block1")
	chain.AddBlock("Block1")
	fmt.Println("Creating block2")
	chain.AddBlock("Block2")

	for _, block := range chain.Blocks {
		fmt.Printf("%s\n", block.Data)
	}
}
