package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Data  []byte
	Prev  []byte
	Hash  []byte
	Nonce int
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

func (block *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(block)
	if err != nil {
		log.Panicln(err)
	}
	return res.Bytes()
}

func Deserialize(input []byte) *Block {
	var res *Block
	decoder := gob.NewDecoder(bytes.NewReader(input))
	err := decoder.Decode(&res)
	if err != nil {
		log.Panicln(err)
	}
	return res
}
func Genesis() *Block {
	return createBlock("Genesis", []byte{})
}
