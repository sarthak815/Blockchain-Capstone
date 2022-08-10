package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Transactions []*Transaction
	Prev         []byte
	Hash         []byte
	Nonce        int
}

func (block *Block) DerieveHash() {
	data := bytes.Join([][]byte{block.Prev, block.HashTransactions()}, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
func createBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{txs, []byte{}, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
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
func Genesis(coinbase *Transaction) *Block {
	return createBlock([]*Transaction{coinbase}, []byte{})
}
