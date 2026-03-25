package main

import(
	"fmt"
)

var firstName string = "Chibueze"

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

	// c = c + "second"
	c += "second"
	fmt.Println(c)

	var b string
	b = "hello"
	var e string
	e = "world"

	fmt.Println(b == e)

	var r string 
	r = "hello"
	var f string
	f = "hello"

	fmt.Println(r == f)

	s := "Hello, world"
	fmt.Println(s)

	n := 5
	fmt.Println(n)

	//Naming variable
	name := "Linda"
	fmt.Println("My dogs name is", name)

	dogsName := "Max" //The variable used camelCase It is also sometimes referred to as mixedCase, BumpyCaps, camelBack, or HumpBack (in some other languages, the first letter is also capitalized).
	fmt.Println("My dogs name is", dogsName)

	fmt.Println(firstName)

	var(
		z = "Chibueze"
		u = 7
		y = "Maxwell"
	)

	fmt.Println(z)
	fmt.Println(u)
	fmt.Println(y)
}

func f()  {
	fmt.Println(firstName)
}