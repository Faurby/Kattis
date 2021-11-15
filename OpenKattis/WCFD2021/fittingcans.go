package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	numOfCans, _ := strconv.Atoi(sc.Text())
	var list []int
	for i := 0; i < numOfCans; i++ {
		sc.Scan()
		can, _ := strconv.Atoi(sc.Text())
		list = append(list, can)
	}
	sort.Ints(list)

}
