package rest

import (
	ohandler "app/infra/rest/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//Route defines a unique route for a REST request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	QueryPairs  []string
}

//Routes define an array of routes supported by the application
type Routes []Route

//NewRouter returns a new Router which routes the REST request to the unique Handler
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var httpHandler http.Handler
		httpHandler = route.HandlerFunc
		httpHandler = ohandler.Logger(httpHandler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(httpHandler).
			Queries(route.QueryPairs...)
	}
	return router
}

//Index returns a welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/v1/",
		HandlerFunc: Index,
	},
	Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/v1/about",
		HandlerFunc: ohandler.About,
	},
}
