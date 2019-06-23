package main

// required package imports
import (
  "net/http"
  "log"
)

// main entry point of application
func main() {

  // call to read config values
  ReadConf()

  // call to get api router object
  router := NewRouter()

  // server start call
	log.Fatal(http.ListenAndServe(configuration.PORT, router))
}
