package client

import (
	"fmt"
	"github.com/swe-zzf/rocket/src/blockchain/core"
	"os"
)

func (cli *CLI) GetBalance(address string) {
	bc := core.GetBlockChainObject()
	if bc == nil {
		fmt.Println("没有BlockChain，无法查询。。")
		os.Exit(1)
	}
	defer bc.DB.Close()
	//total := bc.GetBalance(address,[]*Transaction{})
	utxoSet := &core.UTXOSet{bc}
	total := utxoSet.GetBalance(address)

	fmt.Printf("%s,余额是：%d\n", address, total)
}
