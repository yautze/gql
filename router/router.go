package router

import (
	"gql/middle"

	"github.com/kataras/iris/v12"
)

// SetRouter -
func SetRouter(app *iris.Application) {
	r := app.Party("/api")

	r.Get("/test", func(ctx iris.Context) {
		ctx.JSON(map[string]interface{}{
			"Test": "OK",
		})
	})

	r.Get("/", middle.PlaygroundHandler())
	r.Post("/query", middle.GraphqlHandler())
}
