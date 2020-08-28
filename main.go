package main

import (
	"crypto/sha256"
	"fmt"
)

// Block 区块的数据类型
type Block struct {
	PrevBlockHash []byte //上个区块的哈希
	Hash          []byte //当前区块的哈希
	Data          []byte //数据
}

// SetHash Block下的方法 先来个简单的计算当前区块哈希的函数
func (block *Block) SetHash() {
	var data []byte
	data = append(data, block.PrevBlockHash...)
	data = append(data, block.Data...)
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

// NewBlock 创建区块的工厂函数
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{}, //先给个空的 后续会填充
		Data:          []byte(data),
	}
	block.SetHash()
	return &block
}

// BlockChain 区块链的数据类型
type BlockChain struct {
	Blocks []*Block
}

// AddBlock 向区块链中添加区块
func (bc *BlockChain) AddBlock(data string) {
	// 创建区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(data, prevHash)
	// 添加到区块链中
	bc.Blocks = append(bc.Blocks, block)
}

// NewBlockChain 创建区块链的工厂
func NewBlockChain() *BlockChain {
	// The Times 03/Jan/2009 Chancellor on brink of second bailout for banks
	genesisBlock := NewBlock("创世区块的数据", []byte{0x0000000000000000})
	bc := BlockChain{Blocks: []*Block{genesisBlock}}
	return &bc
}

func main() {
	// fmt.Printf("hello world")
	bc := NewBlockChain()
	bc.AddBlock("新区块1的数据")
	bc.AddBlock("新区块2的数据")
	for i, block := range bc.Blocks {
		fmt.Printf("++++++++++ 区块高度:%d ++++++++++\n", i)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
	}
}
