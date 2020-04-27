package core

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var maxNonce = math.MaxInt64

const targetBits = 24

/**
 * 定义工作量证明
 */
type ProofOfWork struct {
	block *Block
	target *big.Int
}

/**
 * 构建工作量证明
 */
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

/**
 * 准备数据
 */
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

/**
 * 工作量证明算法
 */
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r%x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

/**
 * 验证工作量证明
 */
func (pow *ProofOfWork) Validate() bool {
	var hasInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hasInt.SetBytes(hash[:])

	isValid := hasInt.Cmp(pow.target) == -1

	return isValid
}