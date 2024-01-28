package main

import (
	"fmt"
)

func main() {
	var output int

	// drop your array of number here
	input := []int{1, 5, 8, 4, 5, 9, 3, 0, 4, 6, 2, 7, 34, 6543, 7842, 24}

	for _, val := range input {
		if val > output {
			output = val
		}
	}

	fmt.Println(output)
}
