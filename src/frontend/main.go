// main
package main

import "ppzzl.com/tinyblog-go/frontend/app"

func main() {
	ctx := app.NewContextImpl()
	ctx.GetWebService().Run()
}
