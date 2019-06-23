package main

// required package imports
import (
  "encoding/json"
  "strconv"
)

var currentToDoId int

// function to store todoItem in redis and call msgbroker on successful insertion
func createNewTodoItem(t ToDoItem){
  currentToDoId += 1

  t.Id = currentToDoId

  // acquire redis connection object
  conn := RedisConnect()
  defer conn.Close()

  // marshal Post to JSON blob
  b, err := json.Marshal(t)
  HandleError(err)

  // save JSON to redis
  reply, err := conn.Do("SET", "todo:" + strconv.Itoa(t.Id), b)
  HandleError(err)

  // call msgbroker on successful insertion for intimation of node service
  if reply == "OK" {
    callMsgBroker(currentToDoId)
  }
}
