package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

type Block struct {
	Data       string
	Timestamp  int
	BlockTime  time.Time
	Hash       string
	ParentHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.ParentHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string, timestamp int, blocktime time.Time) *Block {
	newBlock := Block{data, timestamp, blocktime, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	blocktime := time.Now()
	timestamp := blocktime.Unix()
	b.blocks = append(b.blocks, createBlock(data, int(timestamp), blocktime))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
