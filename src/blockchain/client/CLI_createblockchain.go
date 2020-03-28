package client

import (
	"fmt"
	"github.com/swe-zzf/rocket/src/blockchain/core"
	"os"
)

func (cli *CLI) CreateBlockChain(address string) {
	//fmt.Println("创世区块。。。")
	core.CreateBlockChainWithGenesisBlock(address)

	//重置：
	bc := core.GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有数据库。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	utxoSet := &core.UTXOSet{bc}
	utxoSet.ResetUTXOSet()
}
