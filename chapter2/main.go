//TYPES

package main

import (
	"fmt"
)

func main() {
	//Numbers : int, unit, unintpr, float
	fmt.Println("1 + 1 =", 1+1)

	//Strings
	fmt.Println(len("Hello, world"))
	fmt.Println("Hello, world"[1])
	fmt.Println("Hello, " + "world")

	//Booleans : true or false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(!false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Exercise
	fmt.Println("32,132 × 42,452 =", 32132 * 42452 )
	fmt.Println(len("MAXCHICHAR"))
}