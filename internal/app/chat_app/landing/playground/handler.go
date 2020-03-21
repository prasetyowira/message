package playground

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

// GraphiQL is an in-browser IDE for exploring GraphiQL APIs.
// This handler returns GraphiQL when requested.
//
// For more information, see https://github.com/graphql/graphiql.

func RegisterHTTPHandlers(router *mux.Router) {
	router.Path("/playground").Methods(http.MethodGet).Handler(ServeHTTP())
}

func ServeHTTP() http.Handler {
	return playground.Handler("Messaging Graphql", "/graphql", )
}

