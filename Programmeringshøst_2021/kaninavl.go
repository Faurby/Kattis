package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	var fibArray []big.Int
	fibArray = append(fibArray, *big.NewInt(1), *big.NewInt(1))
	fibonacciNumber, _ := strconv.Atoi(sc.Text())
	for i := 0; i < fibonacciNumber; i++ {
		if i > 1 {
			var temp big.Int
			fibArray = append(fibArray, *temp.Add(&fibArray[i-1], &fibArray[i-2]))
		}
		fmt.Println(fibArray[i].String())
	}
}
