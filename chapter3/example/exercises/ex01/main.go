// Question:  Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius (C = (F − 32) * 5/9).

package main

import(
	"fmt"
)

var fahNum float64

func main()  {
	fmt.Println("Enter Fahrenheit number: ")
	fmt.Scanf("%f", &fahNum) // User input. Don't forget the &

	C := (fahNum - 32) * 5/9 //The mechanism 

	fmt.Println("Celsius = ", C)
}