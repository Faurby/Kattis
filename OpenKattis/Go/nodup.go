package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var words []string

	for sc.Scan() {
		word := sc.Text()

		if len(words) == 0 {
			words = append(words, word)
			continue
		}

		sort.Strings(words)
		result := binarySearch(words, word)

		if result != -1 {
			fmt.Print("no")
			return
		} else {
			words = append(words, word)
		}
	}

	fmt.Println("yes")
	return
}

func binarySearch(a []string, search string) (result int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result = binarySearch(a[:mid], search)
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	return
}
