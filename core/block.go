package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

/**
 * 计算哈希
 */
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

/**
 * 创建区块
 */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

/**
 * 生成创世块，也即第一个区块
 */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
