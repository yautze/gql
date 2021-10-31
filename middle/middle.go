package middle

import (
	"gql/controller"
	"gql/controller/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kataras/iris/v12"
)

// GraphqlHandler - defines the gqlgen graphQL server handler
func GraphqlHandler() func(iris.Context) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &controller.Resolver{}}))

	return func(c iris.Context) {
		srv.ServeHTTP(c.ResponseWriter(), c.Request())
	}
}

// PlaygroundHandler - defines the playground handler to expose our playground
func PlaygroundHandler() func(iris.Context) {
	h := playground.Handler("GraphQL playground", "/api/query")
	return func(ctx iris.Context) {
		h.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	}
}
