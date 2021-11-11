package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	word := sc.Text()

	count := make(map[string]int)
	for i := 0; i < len(word); i++ {
		count[string(word[i])]++
	}

	for _, value := range count {
		if value > 1 {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)

}
