package main

import (
	"fmt"
	"strings"
)

func main() {
	var output bool = true

	// drop your text here
	input := "kasur haji ijah rusak"
	arrstr := strings.Split(input, "")

	for idx, val := range arrstr {
		if val != arrstr[len(arrstr)-1-idx] {
			output = false
			break
		}
	}

	fmt.Println(output)
}
