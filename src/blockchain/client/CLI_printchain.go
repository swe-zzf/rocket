package client

import (
	"fmt"
	"github.com/swe-zzf/rocket/src/blockchain/core"
	"os"
)

func (cli *CLI) PrintChains() {
	//cli.BlockChain.PrintChains()
	bc := core.GetBlockChainObject() //bc{Tip,DB}
	if bc == nil {
		fmt.Println("没有BlockChain，无法打印任何数据。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	bc.PrintChains()
}
