package main

import (
	"fmt"
)

func main() {

	// drop your number here
	input := 5

	for i := 1; i <= input; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
