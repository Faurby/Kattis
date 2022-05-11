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

	n, _ := strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		first, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		second, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		third, _ := strconv.Atoi(sc.Text())

		if first+second == third {
			fmt.Println("Possible")
		} else if first-second == third {
			fmt.Println("Possible")
		} else if second-first == third {
			fmt.Println("Possible")
		} else if first*second == third {
			fmt.Println("Possible")
		} else if float64(first)/float64(second) == float64(third) {
			fmt.Println("Possible")
		} else if float64(second)/float64(first) == float64(third) {
			fmt.Println("Possible")
		} else {
			fmt.Println("Impossible")
		}
	}
}
