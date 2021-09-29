package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	maxSheep := 0
	minSheep := 0
	currentSheep := 0
	for sc.Scan() {

		line := sc.Text()

		if line == "Får ind" {
			currentSheep++
			if currentSheep > maxSheep {
				maxSheep = currentSheep
			}
		} else if line == "Får ud" {
			currentSheep--
			if currentSheep < minSheep {
				minSheep = currentSheep
			}

		}

	}

	fmt.Println(Abs(maxSheep - minSheep))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
