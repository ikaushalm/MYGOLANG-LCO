package main

import "fmt"

//here the capital Letter will mark this variable as public
const LoginToken string = "srdtfyguhijokpghbjnk"

func main() {
	//to make function or variable private we use uppercase for the variable name to make it public

	//These are direct assignment

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of the variable is %T \n", isLoggedIn)

	var myName string = "Kaushal"
	fmt.Println(myName)
	fmt.Printf("Type of the variable is %T \n ", myName)

	var myNumber int = 7786868782
	fmt.Println(myNumber)
	fmt.Printf("Type of the variable is %T \n", myNumber)

	var myPercentage float32 = 23.048
	fmt.Println(myPercentage)
	fmt.Printf("Type of the variable is %T \n", myPercentage)

	var number = 78465
	fmt.Println(number)
	fmt.Printf("Type of the variable is %T \n", number)

	//using shorthand operator

	snumber := 78465
	fmt.Println(snumber)
	fmt.Printf("Type of the variable is %T \n", snumber)

	//default values and aliases

	var anotherVariable int

	fmt.Println(anotherVariable)
	fmt.Printf("Type of the variable is %T \n", anotherVariable)

	var anotherstring string

	fmt.Println(anotherstring)
	fmt.Printf("Type of the variable is %T \n", anotherstring)

	//implicit decides the variable type

	var website = "learncodeonlinein"

	fmt.Println(website)
	fmt.Printf("Type of the variable is %T \n", website)

	//most easiest walrus it works under function only
	numberOfUser := 45

	fmt.Println(numberOfUser)

	//acccessing from public
	fmt.Println(LoginToken)

}
