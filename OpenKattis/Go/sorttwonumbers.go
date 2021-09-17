package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a < b {
		fmt.Print(a)
		fmt.Print(" ")
		fmt.Print(b)
	} else {
		fmt.Print(b)
		fmt.Print(" ")
		fmt.Print(a)
	}
}
