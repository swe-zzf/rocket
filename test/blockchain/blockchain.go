package main

import (
	"gustz.com/rocket/v1.0/test/blockchain/core"
	"gustz.com/rocket/v1.0/test/blockchain/rpc"
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
