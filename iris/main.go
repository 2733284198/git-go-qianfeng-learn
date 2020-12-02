package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	// This handler will match /user/john but will not match /user/ or /user
	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		_, _ = ctx.Writef("Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	app.Get("/user/{name}/{action:path}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		action := ctx.Params().Get("action")
		message := name + " is " + action
		_, _ = ctx.WriteString(message)
	})

	_ = app.Listen(":8080")
}
