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
	k, _ := strconv.Atoi(sc.Text())

	var list []int

	for i := 1; i <= n; i++ {
		list = append(list, i)
	}

	if n <= k {
		fmt.Println("impossible")
		return
	}

	for i := 1; i <= k; i++ {
		fmt.Printf("open %d using %d\n", i, i+1)
	}
}
