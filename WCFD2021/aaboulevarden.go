package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type beer struct {
	name  string
	index int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	bars, _ := strconv.Atoi(sc.Text())
	var beers []beer
	for i := 0; i < bars; i++ {
		sc.Scan()
		beers = append(beers, beer{sc.Text(), i})
	}
	sort.Sort(beers())
	sc.Scan()
	route, _ := strconv.Atoi(sc.Text())
	for i := 0; i < route; i++ {
		sc.Scan()
		pos, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		goal := sc.Text()
		if beers[pos] == goal {
			fmt.Println(0)
			continue
		}

		distanceDown := binarySearch()

		ind := 0
		for {
			if pos+ind < bars && beers[pos+ind] == goal {
				fmt.Println(ind)
				break
			} else if pos-ind >= 0 && beers[pos-ind] == goal {
				fmt.Println(-ind)
				break
			}
			ind++
		}
	}
}

func binarySearch(needle string, haystack []string, low int, high int) (int, bool) {

	low = low
	high = high

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return 0, false
	}

	return low, true
}
