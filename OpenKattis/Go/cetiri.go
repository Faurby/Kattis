package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ := strconv.Atoi(sc.Text())

	list := []int{a, b, c}

	sort.Ints(list)

	difference := 900000000
	for i := 1; i < 3; i++ {
		if list[i]-list[i-1] < difference {
			difference = list[i] - list[i-1]
		}
	}

	ind := 0
	for {
		if ind == len(list)-1 {
			fmt.Println(list[ind] + difference)
			break
		}
		if list[ind+1] != list[ind]+difference {
			fmt.Println(list[ind] + difference)
			break
		}
		ind++
	}
}
