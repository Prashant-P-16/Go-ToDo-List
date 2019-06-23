package main

// generic struct for all ToDoItems
type ToDoItem struct{
  Id int
  Title string `json:"title"`
  Desc string `json:"desc"`
}

var ToDoItems []ToDoItem
