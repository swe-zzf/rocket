package rpc

import (
	"encoding/json"
	"gustz.com/rocket/test/blockchain/core"
	"io"
	"log"
	"net/http"
)

/**
 * RPC服务
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-24
 */
func readApi(bc *core.BlockChain) {
	http.HandleFunc("/block_chain/read", func(writer http.ResponseWriter, request *http.Request) {
		var data []byte
		var err error
		if data, err = json.MarshalIndent(bc.GetBlocks(), "", " "); err != nil {
			log.Println("to json error=", err)
			return
		}
		returnResp(writer, string(data))
	})
}

func writeApi(bc *core.BlockChain) {
	http.HandleFunc("/block_chain/write", func(writer http.ResponseWriter, request *http.Request) {
		bc.Send(request.FormValue("data"))
		returnResp(writer, "success")
	})
}

func returnResp(writer http.ResponseWriter, data string) {
	if n, err := io.WriteString(writer, string(data)); err != nil {
		log.Println("n=", n, ",to write error=", err)
		return
	}
}

func RunServer(bc *core.BlockChain) {
	readApi(bc)
	writeApi(bc)
	err := http.ListenAndServe(":8888", nil)
	log.Fatal(err)
}
