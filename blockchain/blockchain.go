package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Block struct {
	Data       string    `json:"data"`
	Timestamp  int       `json:"timestamp"`
	BlockTime  time.Time `json:"blockTime"`
	Hash       string    `json:"hash"`
	ParentHash string    `json:"parentHash,omitempty"`
	Height     int       `json:"height"`
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
	newBlock := Block{data, timestamp, blocktime, "", getLastHash(), len(GetBlockchain().blocks) + 1}
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

var ErrNotFound = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}

	return b.blocks[height-1], nil
}
