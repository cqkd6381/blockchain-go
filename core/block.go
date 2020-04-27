package core

import (
	"time"
)

/**
 * 定义区块
 * Timestamp 当前时间戳，也就是区块创建的时间
 * Data 区块存储的实际有效信息，也就是交易
 * PrevBlockHash 前一个区块的哈希，即父哈希
 * Hash 当前块的哈希，用户校验区块的安全性
 */
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

/**
 * 创建区块
 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

/**
 * 生成创世块，也即第一个区块
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
