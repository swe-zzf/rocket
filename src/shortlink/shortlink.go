package main

import (
	"gustz.com/rocket/v1.0/src/shortlink/api"
	"gustz.com/rocket/v1.0/src/shortlink/core"
)

/**
 * XXX 
 *
 * @author gavin.z, swe.zzf@gmail.com
 * @since  2020-03-26
 */
func main() {
	a := api.App{}
	a.Initialize(core.GetEnv())
	a.Run(":8888")
}
