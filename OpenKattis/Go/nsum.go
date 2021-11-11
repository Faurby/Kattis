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

	var sum int
	for i := 0; i < n; i++ {
		sc.Scan()
		number, _ := strconv.Atoi(sc.Text())
		sum += number
	}
	fmt.Println(sum)
}
