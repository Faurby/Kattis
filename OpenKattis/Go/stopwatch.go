package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	if N%2 != 0 {
		fmt.Println("still running")
	} else {
		sum := 0
		for i := 0; i < N; i += 2 {
			var a, b int
			fmt.Scan(&a, &b)
			sum += b - a
		}
		fmt.Println(sum)
	}
}
