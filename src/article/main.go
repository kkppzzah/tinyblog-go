// main
package main

import "ppzzl.com/tinyblog-go/article/app"

func main() {
	ctx := app.NewContextImpl()
	ctx.GetGRPCServer().Run()
}
