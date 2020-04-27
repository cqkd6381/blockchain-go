package core

/**
 * 定义区块链
 */
type Blockchain struct {
	Blocks []*Block
}

/**
 * 添加新的区块到区块链中
 */
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

/**
 * 创建包含创世块的区块链
 */
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
