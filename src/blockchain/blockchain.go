package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	Blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(Getblockchain().Blocks)
	if totalBlocks == 0 {
		return ""
	}
	return Getblockchain().Blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) Addblock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
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

func (b *blockchain) Allblocks() []*Block {
	return b.Blocks
}
