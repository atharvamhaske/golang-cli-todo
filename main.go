package main 

import "fmt"

func main() {
    todos := Todos{}
    todos.add("Complete Go ")
    todos.add("Complete Devops")
    fmt.Printf("%+v\n\n", todos)
    todos.delete(0)
    fmt.Printf("%+v", todos)
} 