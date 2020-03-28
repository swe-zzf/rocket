package main

import (
	"github.com/swe-zzf/rocket/test/blockchain/core"
	"github.com/swe-zzf/rocket/test/blockchain/rpc"
)

/**
 * 测试函数
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-24
 */
func main() {
	bc := core.NewBlockChain()
	bc.Send("one data")
	bc.Send("two data")
	bc.Send("three data")
	bc.Print()
	rpc.RunServer(bc)

}
