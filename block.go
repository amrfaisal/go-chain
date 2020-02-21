package main

import (
	"bytes"
	"crypto/sha1"
	"strconv"
	"time"
)

//Block
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

//SetHash
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha1.Sum(headers)

	b.Hash = hash[:]
}

//NewBlock
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

//NewGenesisBlock
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
