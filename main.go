package main

func main() {
	todos := Todos{}
	todos.add("Complete Go")
	todos.add("Complete Devops")

	// Mark first todo as completed
	todos.toggle(0)

	todos.print()
}
