package main 

import "fmt"

func slicesDemo() {
	numbers := []int {10, 20, 30, 40, 50}
	indexToDelete := 2 

    // get everything before number 2 and after 2 and make slices and append them later
    before := numbers[:indexToDelete]

	after := numbers[indexToDelete+1:]

	// to join both slices using append

	numbers = append(before, after...)

	fmt.Println(numbers)  //Output: [10 20 40 50 ]


 }

// alternate one liner method to do in go idiom
// numbers = append(numbers[:indexToDelete], numbers[indexToDelete+1:]...)