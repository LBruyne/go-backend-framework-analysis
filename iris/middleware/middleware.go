package middleware

import "github.com/kataras/iris/v12"

func SetMiddleware() *iris.Application {
	app := iris.New()
	// or app.Use(before) and app.Done(after).
	app.Get("/", middleware)
	app.Run(iris.Addr(":8080"))

	app.UseGlobal(before)
	app.DoneGlobal(after)
	return app
}

func before(ctx iris.Context) {
	shareInformation := "this is a sharable information between handlers"

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")
}

func middleware(ctx iris.Context) {
	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}
