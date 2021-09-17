package main

import (
	"fmt"
	"strconv"
)

func main() {
	var antal int
	fmt.Scan(&antal)

	var fishy bool

	for i := 1; i <= antal; i++ {
		var input string
		fmt.Scan(&input)

		if input != "mumble" {
			input, err := strconv.Atoi(input)
			_ = err
			if i != input {
				fmt.Println("something is fishy")
				fishy = true
				break
			}
		}
	}

	if !fishy {
		fmt.Println("makes sense")
	}

}
