// Question: Write another program that converts from feet into meters (1 feet = 0.3048 m)

package main

import(
	"fmt"
)

var feet float64

func main()  {
	fmt.Println("Enter Feet number: ")
	fmt.Scanf("%f", &feet) //User input

	meter := feet * 0.3048 //The mechanism i.e formula

	fmt.Println("Meter = ", meter)
}