package main

import "fmt"

func main() {
	var input int

	fmt.Scan(&input)

	var sum int
	for i := 0; i < input; i++ {
		var rod int
		fmt.Scan(&rod)

		sum += rod
	}

	fmt.Println(sum - (input - 1))
}
