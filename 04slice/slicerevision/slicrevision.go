package main

import "fmt"

func main() {

	//slice

	cards := []string{"abcd", "efgh", "ijkl"}

	// range mean to iterate over slice
	//:= to assign values-- surprising

	// with for loop after every loop we are throughing the last variables that is i and value
	//inside body we are printing those values

	for i, value := range cards {
		fmt.Println(i, value)
	}

}
