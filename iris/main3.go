package main

import (
	"github.com/kataras/iris"
	//"github.com/kataras/iris/core/router" 1.9之后需要引入
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {})
	app.Run(iris.Addr(":8080"))
	app.PartyFunc("/cpanel", func(child iris.Party) {
		child.Get("/", func(ctx iris.Context) {})
	})
	// OR
	cpanel := app.Party("/cpanel")
	cpanel.Get("/", func(ctx iris.Context) {})
}
