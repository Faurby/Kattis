package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	input := sc.Text()

	r, _ := regexp.Compile("(:\\)|;-\\)|:-\\)|;\\))")

	results := r.FindAllStringIndex(input, -1)

	for _, v := range results {
		fmt.Println(v[0])
	}
}
