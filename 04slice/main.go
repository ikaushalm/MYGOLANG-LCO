package main

import "fmt"

func main() {

	//slice is an array of flexible length similar to arraylist but it is statically types it is an array with fixed length

	cards := []string{GetString(), "Testing Slice"}

	for i, val := range cards {
		fmt.Println(i, val)

	}

}

func GetString() string {
	return "ABCD"
}
