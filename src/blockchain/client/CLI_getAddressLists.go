package client

import (
	"fmt"
	"github.com/swe-zzf/rocket/src/blockchain/wallet"
)

func (cli *CLI) GetAddressLists() {
	fmt.Println("钱包地址列表为:")
	//获取钱包的集合，遍历，依次输出
	wallets := wallet.GetWallets()
	for address, _ := range wallets.WalletMap {

		fmt.Printf("\t%s\n", address)
	}
}
