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
	sum := 0.0
	counter := 0.0
	for i := 0; i < n; i++ {
		sc.Scan()
		bat, _ := strconv.Atoi(sc.Text())
		if bat != -1 {
			sum += float64(bat)
			counter++
		}
	}
	fmt.Println(float64(sum / counter))
}
