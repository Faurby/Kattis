package main

import "fmt"

func main() {
	n := float32(1)
	k := float32(1)

	var sum float32
	for i := 0; i < int(k); i++ {
		sum += float32(1)
	}

	overall := float32(n - k)

	output1 := float32((sum + (overall * 3)) / n)
	output2 := float32((sum + (overall * -3)) / n)

	fmt.Printf("%f %f", output2, output1)

}
