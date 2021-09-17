package main

import "fmt"

func main() {
	var A, B, C, I, J, K float32
	fmt.Scanln(&A, &B, &C)
	fmt.Scanln(&I, &J, &K)

	var drinksToMake float32 = A / I

	if drinksToMake > (B / J) {
		drinksToMake = B / J
	}
	if drinksToMake > (C / K) {
		drinksToMake = C / K
	}

	fmt.Printf("%f %f %f", A-drinksToMake*I, B-drinksToMake*J, C-drinksToMake*K)

}
