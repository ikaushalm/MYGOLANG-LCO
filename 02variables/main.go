package main

import "fmt"

func main() {
	//to make function or variable private we use uppercase for the variable name to make it public

	//These are direct assignment

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of the variable is%T \n", isLoggedIn)

	var myName string = "Kaushal"
	fmt.Println(myName)
	fmt.Printf("Type of the variable is%T \n ", myName)

	var myNumber int = 7786868782
	fmt.Println(myNumber)
	fmt.Printf("Type of the variable is%T \n", myNumber)

	var myPercentage float32 = 23.048
	fmt.Println(myPercentage)
	fmt.Printf("Type of the variable is%T \n", myPercentage)

	var number = 78465
	fmt.Println(number)
	fmt.Printf("Type of the variable is%T \n", number)

	//using shorthand operator

	snumber := 78465
	fmt.Println(snumber)
	fmt.Printf("Type of the variable is%T \n", snumber)

}
