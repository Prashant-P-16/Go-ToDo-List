package main

// required package imports
import (
	"github.com/gorilla/mux"
)

// function to loop through all the defined routes and return api router object for all of them
func NewRouter() *mux.Router {
		// generate mux api router object
  	router := mux.NewRouter().StrictSlash(true)

		// loop through all the defined routes and assign handler function, method and path
		// to the respective router object
  	for _, route := range routes {
  	router.HandleFunc(route.Path, route.Handle).Methods(route.Method)
  }
	return router
}
