package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
 * HTTP服务
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-24
 */
func indexApi() {
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		user := request.FormValue("user")
		n, err := fmt.Fprintln(writer, "欢迎 "+user+" 访问index！")
		log.Println("n=", n, ",error=", err)
	})
}

func loginApi() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		user := request.FormValue("user")
		n, err := fmt.Fprintln(writer, "欢迎 "+user+" 登录成功！")
		log.Println("n=", n, ",error=", err)
	})
}

func main() {
	indexApi()
	loginApi()
	err := http.ListenAndServe(":8888", nil)
	fmt.Println("error=", err)
	time.Sleep(60)
}
