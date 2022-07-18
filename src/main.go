package main

import (
	"fmt"
	"go-mod/blockchain"
)

func main() {
	chain := blockchain.Getblockchain()
	chain.Addblock("Second block")
	chain.Addblock("Third block")
	chain.Addblock("Fourth block")
	for _, block := range chain.Allblocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
