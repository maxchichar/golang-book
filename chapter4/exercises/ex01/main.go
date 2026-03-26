/*
What does the following program print?

i := 10
if i > 10 {
 fmt.Println("Big")
} else {
 fmt.Println("Small")
}
*/

package main

import(
	"fmt"
)

func main()  {
	fmt.Print("Input a number to see if it's bigger than 10: ")
	var i int
	fmt.Scanln(&i)
	fmt.Println()

	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}