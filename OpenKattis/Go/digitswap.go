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

	fmt.Printf("%s%s", string(input[1]), string(input[0]))
}
