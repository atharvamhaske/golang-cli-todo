package main

import "fmt"

func conceptsDemo() {
	x := 100
	// a normal variable holding a value

	p := &x
	// this is a pointer "p" holding the memory address of "x"
	// here p itself does not hold value 100, it hold the location of "x"

	value := *p

	// this is called dereferencing the pointer "p"
	// the * operator here means "get the value at this address"

	fmt.Println("Value of x:", x)                   // Prints 100
	fmt.Println("Memory address in p:", p)          // Prints a memory address like 0xc0000180a8
	fmt.Println("Value by dereferencing p:", value) // Prints 100

	// we can change the value at the address where p is pointing to.
	*p = 250

	fmt.Println("New value of x:", x) // Prints 250
}

func pointerExample() {
	city := "Pune"
	cityPointer := &city

	fmt.Println("Before change, city is:", city)

	// Use the pointer to change the value at its address.
	*cityPointer = "Mumbai"

	fmt.Println("After change, city is:", city)
}

//we are making toggle method to change task status if completed and if not

// here Todo is the type, Think of it as the blueprint for your list.
//*Todo means a pointer to a Todos list, a pointer is just memory address which tells the function where to find the original list.
// also todos: this is just a variable name it is the name used inside the method to refer to that pointer

// also t := (*todos) here todos is a pointer ( a memory address) and *todos is a value which we get at that address just like saying that we are dereferncing the pointer. this line gets the actual list of todos

// isCompleted := t[index].Completed , t[index] gets the specific todo from the list.

// .Completed accesses its Completed field, which is either true or false.

// This true or false value is stored in a temporary variable called isCompleted.


//if else 

func pract() {
	age := 19

	if age >= 19 {
		fmt.Println("you are an adult now")
	} else {
		fmt.Println("you are not adult yet")
	}
}