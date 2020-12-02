package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.Get("/hello", func(context iris.Context) {
		_, _ = context.WriteString("hello")
	})

	_ = app.Listen(":8080")
}
