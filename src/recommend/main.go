// main
package main

import "ppzzl.com/tinyblog-go/recommend/app"

func main() {
	ctx := app.NewContextImpl()
	ctx.GetGRPCServer().Run()
}
