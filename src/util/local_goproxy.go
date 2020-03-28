package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goproxy/goproxy"
)

/**
 * local goproxy server
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-28
 */
func RunLoGoProxy() {
	proxy := goproxy.New()
	proxy.GoBinEnv = append(
		os.Environ(),
		"GOPROXY=https://goproxy.cn,direct", // 使用 goproxy.cn 作为上游代理
		"GOPRIVATE=github.com",              // 解决私有模块的拉取问题（比如可以配置成公司内部的代码源）
	)
	fmt.Println("Run local goproxy done!!!")
	if err := http.ListenAndServe(":9999", proxy); err != nil {
		fmt.Println("Run local goproxy err=", err)
		return
	}

}
