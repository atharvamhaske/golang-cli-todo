package main 

import "fmt"

func conceptsDemo () {
	x := 100
	// a normal variable holding a value 

	p := &x
	// this is a pointer "p" holding the memory address of "x"
	// here p itself does not hold value 100, it hold the location of "x"

	value := *p

	// this is called dereferencing the pointer "p"
	// the * operator here means "get the value at this address"
    

    fmt.Println("Value of x:", x)               // Prints 100
    fmt.Println("Memory address in p:", p)    // Prints a memory address like 0xc0000180a8
    fmt.Println("Value by dereferencing p:", value) // Prints 100

	// we can change the value at the address where p is pointing to.
	*p = 250

	fmt.Println("New value of x:", x) // Prints 250
}

package main

import "fmt"

func main() {
    city := "Pune"
    cityPointer := &city

    fmt.Println("Before change, city is:", city)

    // Use the pointer to change the value at its address.
    *cityPointer = "Mumbai"

    fmt.Println("After change, city is:", city)
}