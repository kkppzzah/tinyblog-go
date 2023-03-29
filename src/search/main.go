// main
package main

import "ppzzl.com/tinyblog-go/search/app"

func main() {
	ctx := app.NewContextImpl()
	ctx.GetGRPCServer().Run()
}
