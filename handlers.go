package main

// required package imports
import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

// function to handle request, parse request body & store items in db
func TodoCreate(w http.ResponseWriter, r *http.Request) {
  // get the body of POST request
  reqBody, _ := ioutil.ReadAll(r.Body)

	// unmarshal the request into a new ToDoItem struct
  var todoItem ToDoItem
  json.Unmarshal(reqBody, &todoItem)

	// call to create new todoItem in Redis
  createNewTodoItem(todoItem)
}
