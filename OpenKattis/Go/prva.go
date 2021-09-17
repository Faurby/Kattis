package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var R, C int
	fmt.Scanln(&R, &C)

	board := make([][]string, R)

	for i := 0; i < R; i++ {
		var input string
		fmt.Scan(&input)
		split := strings.Split(input, "")
		board[i] = split
	}

	var words []string
	for i := 0; i < R; i++ {
		var currentList []string
		for j := 0; j < C; j++ {
			current := board[i][j]

			if current == "#" {
				if len(currentList) > 1 {
					words = append(words, strings.Join(currentList, ""))
				}
				currentList = nil
			} else {
				currentList = append(currentList, current)
			}

		}
		if len(currentList) > 1 {
			words = append(words, strings.Join(currentList, ""))
		}
		currentList = nil
	}

	for i := 0; i < C; i++ {
		var currentList []string
		for j := 0; j < R; j++ {
			current := board[j][i]

			if current == "#" {
				if len(currentList) > 1 {
					words = append(words, strings.Join(currentList, ""))
				}
				currentList = nil
			} else {
				currentList = append(currentList, current)
			}

		}
		if len(currentList) > 1 {
			words = append(words, strings.Join(currentList, ""))
		}
		currentList = nil
	}

	sort.Strings(words)

	if len(words) != 0 {
		fmt.Println(words[0])
	}
}
