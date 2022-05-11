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
	input := sc.Text()

	index := 1

	for _, v := range input {
		switch string(v) {
		case "A":
			if index == 1 {
				index = 2
			} else if index == 2 {
				index = 1
			}
		case "B":
			if index == 2 {
				index = 3
			} else if index == 3 {
				index = 2
			}
		case "C":
			if index == 1 {
				index = 3
			} else if index == 3 {
				index = 1
			}
		}
	}
	fmt.Println(index)
}
