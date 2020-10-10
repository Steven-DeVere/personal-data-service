package graphql

import "github.com/tkww/guest-data-service/graphql/generated"

type Resolver struct{}

// NewGraphQLHandler returns an http.Handler that will serve the GraphQL data.
func NewGraphQLHandler() *handler.Server {
	return handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{Env: env, DB: DB}}))
}
