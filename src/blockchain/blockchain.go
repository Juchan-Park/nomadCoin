package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(Getblockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return Getblockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) Addblock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func Getblockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.Addblock("Genesis block")

		})
	}
	return b

}

func (b *blockchain) Allblocks() []*block {
	return b.blocks
}
