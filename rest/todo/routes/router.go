package todo

import (
	"net/http"

	"rest/loggers"

	"github.com/gorilla/mux"
)

func TodoRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = loggers.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
