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

	sc.Scan()
	t, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	fmt.Println(n * t * m)
}
