package client

import (
	"fmt"
	"github.com/swe-zzf/rocket/src/blockchain/core"
	"os"
)

func (cli *CLI) Send(from, to, amount []string) {
	bc := core.GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法转账。。")
		os.Exit(1)
	}
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount)
	//添加更新
	utsoSet := &core.UTXOSet{bc}
	utsoSet.Update()
}
