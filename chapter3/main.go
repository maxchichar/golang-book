package main

import(
	"fmt"
)

func main()  {
	var x string = "Hello, World"
	fmt.Println(x)

	var a string
	a = "King charles"
	fmt.Println(a)

	var c string
	c = "first"
	fmt.Println(c)

	// c = "second"
	// fmt.Println(c)

	c = c + " " + "second"
	fmt.Println(c)
}