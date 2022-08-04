package core

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Data  []byte
	Prev  []byte
	Hash  []byte
	Nonce int
}

type BlockChain struct {
	Blocks []Block
}

func (block *Block) DerieveHash() {
	data := bytes.Join([][]byte{block.Prev, block.Data}, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte(data), prevHash, []byte{}, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
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
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := createBlock(data, prevBlock.Hash)
	block.DerieveHash()
	chain.Blocks = append(chain.Blocks, *block)
}
