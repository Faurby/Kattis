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

	a, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())

	fmt.Println(a + b)
}
