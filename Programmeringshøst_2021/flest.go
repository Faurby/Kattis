package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()

	//arrange words into map. String -> int
	animals := strings.Fields(sc.Text())
	var table = make(map[string]int)

	for _, i := range animals {
		table[i] = 0
	}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {

		sc.Scan()
		line := sc.Text()

		for key, _ := range table {
			count := strings.Count(line, key)
			table[key] += count
		}
	}

	fmt.Println(table)

	greatest := 0
	for _, i := range table {
		if i > greatest {
			greatest = i
		}
	}

	fmt.Println(greatest)

}
