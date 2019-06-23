package main

// required package imports
import (
	 "net/http"
)

// generic struct for all http requests: endpoints, handler, http method
type Route struct {
	Path     		string
	Handle	    func(http.ResponseWriter, *http.Request)
	Method      string
}

type Routes []Route

// define all the routes to be assigned to router
var routes = Routes{
	Route{
		"/todoitem",
		TodoCreate,
		"POST",
	},
}
