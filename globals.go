package main

// required package imports
import (
  "log"
  "os"
  "encoding/json"
)

// generic struct for all Config keys
type Configuration struct {
  PORT              string
  REDIS_NETWORK     string
  REDIS_ADDR        string
  RABBIT_CONN_URL   string
  RABBIT_QNAME      string
}

var configuration Configuration

// function to read and parse config values into a structure object
func ReadConf() {

  // open config file
  file, _ := os.Open("conf.json")
  defer file.Close()

  // decode config json & store in config structure object
  decoder := json.NewDecoder(file)
  err := decoder.Decode(&configuration)
  if err != nil {
    log.Fatalf("Error while reading config: %s", err)
  }
}

// generic function to handle errors
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
