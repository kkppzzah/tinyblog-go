// main
package main

import "ppzzl.com/tinyblog-go/user/app"

func main() {
	ctx := app.NewContextImpl()
	ctx.GetGRPCServer().Run()
}
