package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	bc.AddBlock("Data 1")
	bc.AddBlock("Data 2")
	bc.AddBlock("data 3")

	bci := bc.Iterator()

	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
