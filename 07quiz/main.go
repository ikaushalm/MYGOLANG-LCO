package main

import "fmt"

type book string

type laptopSize float64

//we never use self or this so dont use this
func (this laptopSize) getSizeOfLaptop() laptopSize {
	return this
}

var km laptopSize = 45.00

func (b book) printTitle() {
	fmt.Println(b)
}

func main() {
	var b book = "Harry Potter"
	b.printTitle()
	fmt.Print(km.getSizeOfLaptop())

}
