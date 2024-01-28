package main

import (
	"fmt"
)

func main() {
	var output int = 1

	// drop your number here
	input := 5

	for i := 1; i <= input; i++ {
		output *= i
	}

	fmt.Println(output)
}
