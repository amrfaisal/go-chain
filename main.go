package main

import "fmt"

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Data 1")
	bc.AddBlock("Data 2")
	bc.AddBlock("data 3")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
