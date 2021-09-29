package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	width, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	pieces, _ := strconv.Atoi(scanner.Text())

	var area int
	for i := 0; i < pieces; i++ {

		scanner.Scan()
		width, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		height, _ := strconv.Atoi(scanner.Text())

		area += width * height
	}

	fmt.Println(area / width)
}
