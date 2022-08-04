package Blockchain_Capstone

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data []byte
	Prev []byte
	Hash []byte
}

type BlockChain struct {
	block []Block
}

func (block *Block) DerieveHash() {
	data := bytes.Join([][]byte{block.Prev, block.Data}, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), prevHash, []byte{}}
	block.DerieveHash()
	return block
}

func GenesisBlock() *Block {
	block := createBlock("Genesis", []byte{})
	block.DerieveHash()
	return block

}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]Block{*GenesisBlock()}}
}
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.block[len(chain.block)-1]
	block := createBlock(data, prevBlock.Hash)
	block.DerieveHash()
	chain.block = append(chain.block, *block)
}

func main() {
	fmt.Print("Creating Genesis Block")
	fmt.Println("Creating Block1")
	fmt.Println("Creating block2")
	chain := InitBlockChain()
	chain.AddBlock("Block1")
	chain.AddBlock("Block2")
}
