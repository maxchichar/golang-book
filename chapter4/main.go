// Control Structures

package main

import(
	"fmt"
)

func main()  {
	// For statement

	/*
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
	*/

	// Or

	/*
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	*/ 

	// If satement
	/*
	if i % 2 == 0 {
		// even
	} else {
		// odd
	}

	if i % 2 == 0 {
		// divisible by 2
	} else if i % 3 == 0{
		// divisible by 3
	} else if i % 4 == 0 {
		// divisible by 4
	}
	*/

	// For and If together
	
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}


	// The Switch Statement

	/*
	if i == 0 {
		fmt.Println("zero")
	} else if i == 1 {
		fmt.Println("one")
	} else if i == 2 { 
		fmt.Println("two")
	} else if i == 3 {
		fmt.Println("three")
	} else if i == 4 {
		fmt.Println("four")
	} else if i == 5 {
		fmt.Println("five")
	}
	*/

	// using the if statement is going to be tedious that why Go created the Switch statement

	/*
	switch i {
	case 0: fmt.Println("zero")
	case 1: fmt.Println("one")
	case 2: fmt.Println("two")
	case 3: fmt.Println("three")
	case 4: fmt.Println("four")
	case 5: fmt.Println("five")
	default: fmt.Println("unknown number")
	}
	*/
}