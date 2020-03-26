package core

import (
	"fmt"
	"log"
)

/**
 * 区块链实体操作
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-24
 */
type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	// 创建创世区块
	bc.appendBlock(newFirstBlock())
	return bc
}

func (bc *BlockChain) GetBlocks() []*Block {
	return bc.blocks
}

func (bc *BlockChain) Send(data string) {
	if data == "" {
		log.Println("invalid data.")
		return
	}
	prevBlock := bc.blocks[len(bc.blocks)-1]
	prevBlockSn := prevBlock.Sn
	prevBlockHash := prevBlock.CurrHash
	// 创建普通区块
	newBlock := newBlock(prevBlockSn, prevBlockHash, data)
	// 追加到数组
	bc.appendBlock(newBlock)
}

func (bc *BlockChain) appendBlock(newBlock *Block) {
	if len(bc.blocks) == 0 {
		//log.Println("create first block.")
		bc.blocks = append(bc.blocks, newBlock)
		return
	}
	oldBlock := bc.blocks[len(bc.blocks)-1]
	if isInvalid(*newBlock, *oldBlock) {
		log.Println("invalid block.")
		return
	}
	bc.blocks = append(bc.blocks, newBlock)
}

func isInvalid(newBlock Block, oldBlock Block) bool {
	// s1: 数据
	if newBlock.Data == "" {
		return true
	}
	// 序号
	if (newBlock.Sn - 1) != oldBlock.Sn {
		return true
	}
	// PrevHash值
	if newBlock.PrevHash != oldBlock.CurrHash {
		return true
	}
	// s99: CurrHash值
	currHash := createHash(newBlock)
	if newBlock.CurrHash != currHash {
		return true
	}
	return false
}

func (bc *BlockChain) Print() {
	fmt.Println("")
	for _, block := range bc.blocks {
		fmt.Println("sn:", block.Sn)
		fmt.Println("timestamp:", block.Timestamp)
		fmt.Println("prevHash:", block.PrevHash)
		fmt.Println("currHash:", block.CurrHash)
		fmt.Println("data:", block.Data)
		fmt.Println("")
	}
}
