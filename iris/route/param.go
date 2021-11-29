package route

import (
	"github.com/kataras/iris/v12"
	"regexp"
)

func getParamWithRegisterFunc() *iris.Application {
	app := iris.New()

	// internal validate function
	app.Get("/profile/{name:alphabetical max(255)}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(name)
		// len(name) <=255 otherwise this route will fire 404 Not Found
		// and this handler will not be executed at all.
	})

	// self defined validate function
	latLonExpr := "^-?[0-9]{1,3}(?:\\.[0-9]{1,10})?$"
	latLonRegex, _ := regexp.Compile(latLonExpr)

	// Register your custom argument-less macro function to the :string param type.
	// MatchString is a type of func(string) bool, so we use it as it is.
	app.Macros().Get("string").RegisterFunc("coordinate", latLonRegex.MatchString)

	app.Get("/coordinates/{lat:string coordinate()}/{lon:string coordinate()}",
		func(ctx iris.Context) {
			ctx.Writef("Lat: %s | Lon: %s", ctx.Params().Get("lat"), ctx.Params().Get("lon"))
		})

	return app
}

func getParamFromForm() *iris.Application {
	// POST /post?id=1234&page=1 HTTP/1.1
	// Content-Type: application/x-www-form-urlencoded
	//
	// name=manu&message=this_is_great

	app := iris.Default()
	app.Post("/post", func(ctx iris.Context) {
		id := ctx.URLParam("id")
		page := ctx.URLParamDefault("page", "0")
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		app.Logger().Infof("id: %s; page: %s; name: %s; message: %s",
			id, page, name, message)
	})

	return app
}
