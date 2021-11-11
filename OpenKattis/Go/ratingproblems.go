package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	n, _ := strconv.ParseFloat(sc.Text(), 32)

	sc.Scan()
	k, _ := strconv.ParseFloat(sc.Text(), 32)

	var sum float64
	for i := 0; i < int(k); i++ {
		sc.Scan()
		number, _ := strconv.ParseFloat(sc.Text(), 32)
		sum += number
	}

	overall := n - k

	bestcase := overall * float64(3)
	worstcase := overall * float64(-3)

	output1 := (sum + bestcase) / n
	output2 := (sum + worstcase) / n

	fmt.Printf("%f %f", output2, output1)

}
