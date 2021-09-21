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

	current, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	desired, _ := strconv.Atoi(sc.Text())

	if desired < current {
		result := current - desired
		if result < 180 {
			result *= -1

		} else {
			result = 360 - result
		}

		fmt.Print(result)
	} else if desired == current {
		fmt.Print(0)
	} else {
		result := desired - current
		if result > 180 {
			result = (360 - result) * (-1)
		}
		fmt.Print(result)
	}
}
