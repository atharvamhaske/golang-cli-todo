package main 



func main() {
    todos := Todos{}
    todos.add("Complete Go ")
    todos.add("Complete Devops")
    todos.toggle(0)
    todos.print()
} 