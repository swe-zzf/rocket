package core

import (
	"crypto/sha512"
	"encoding/hex"
	"time"
)

/**
 * 区块实体操作
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-24
 */
type Block struct {
	Sn        int64  // 序号
	Timestamp int64  // 时间戳
	PrevHash  string // 父级Hash值
	CurrHash  string // 当前Hash值
	Data      string // 数据
}

/**
 * 定义block的方法
 */
// 创建区块
func newBlock(prevBlockSn int64, prevBlockHash string, data string) *Block {
	block := Block{}
	block.Sn = prevBlockSn + 1
	block.Timestamp = time.Now().UnixNano()
	block.PrevHash = prevBlockHash
	// s1:
	block.Data = data
	// s2:
	block.CurrHash = createHash(block)
	return &block
}

// 创建创世区块
func newFirstBlock() *Block {
	block := Block{}
	block.Sn = -1
	block.PrevHash = ""
	// s1:
	block.Data = "First data"
	// s2:
	block.CurrHash = createHash(block)
	return newBlock(block.Sn, block.CurrHash, block.Data)
}

// 生成Hash值
func createHash(block Block) string {
	// Hash值=hash(序号+时间戳+父级Hash值+数据)
	value := string(block.Sn) + string(block.Timestamp) + block.PrevHash + block.Data
	rs := sha512.Sum512([]byte(value))
	return hex.EncodeToString(rs[:])
}
