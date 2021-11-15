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
	p, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	s, _ := strconv.Atoi(sc.Text())

	for i := 1; i <= s; i++ {
		if i%p == 0 && i%q == 0 {
			fmt.Println("yes")
			return
		}
	}
	fmt.Println("no")

}
