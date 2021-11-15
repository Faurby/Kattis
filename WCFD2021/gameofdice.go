package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	var lines []string
	sc.Scan()
	for i := 0; i < 3; i++ {
		sc.Scan()
		lines = append(lines, sc.Text())
	}
	var listDice = make([]int, 5)
	for i := 0; i < 3; i++ {
		for ind, v := range strings.Split(lines[i], "| |") {
			count := strings.Count(v, "o")
			listDice[ind] += count
		}
	}

	fmt.Print(countSingles(listDice, 1), countSingles(listDice, 2), countSingles(listDice, 3), countSingles(listDice, 4), countSingles(listDice, 5), countSingles(listDice, 6))
	fmt.Println()
	fmt.Print(countDoubles(listDice), countTwoPair(listDice), threeofakind(listDice), fourofakind(listDice), fullhouse(listDice), smallstraight(listDice), largestraight(listDice), yatzy(listDice), chance(listDice))
}

func countNumber(input []int, countNumber int) int {
	sum := 0
	for _, v := range input {
		if v == countNumber {
			sum++
		}
	}
	return sum
}

func countSingles(input []int, number int) int {

	return countNumber(input, number) * number
}

func countDoubles(input []int) int {
	sort.Ints(input)
	if input[4] == input[3] {
		return 2 * input[4]
	}
	if input[3] == input[2] {
		return 2 * input[3]
	}
	if input[2] == input[1] {
		return 2 * input[2]
	}
	if input[1] == input[0] {
		return 2 * input[1]
	}
	return 0
}

func countTwoPair(input []int) int {
	sort.Ints(input)
	var firstpair int
	var points int
	var pairs int
	for i := 1; i < 5; i++ {
		if input[i] == input[i-1] && (firstpair != input[i]) {
			points += input[i] * 2
			pairs++
			firstpair = input[i]
			i++
		}
	}
	if pairs == 2 {
		return points
	}
	return 0
}

func threeofakind(input []int) int {
	sort.Ints(input)
	for i := 2; i < 5; i++ {
		if input[i] == input[i-1] && input[i] == input[i-2] {
			return 3 * input[i]
		}
	}
	return 0
}

func fourofakind(input []int) int {
	sort.Ints(input)
	for i := 3; i < 5; i++ {
		if input[i] == input[i-1] && input[i] == input[i-2] && input[i] == input[i-3] {
			return 4 * input[i]
		}
	}
	return 0
}

func yatzy(input []int) int {
	for i := 1; i < 5; i++ {
		if input[i] != input[i-1] {
			return 0
		}
	}
	return 50
}

func smallstraight(input []int) int {
	sort.Ints(input)
	if input[0] == 1 && input[1] == 2 && input[2] == 3 && input[3] == 4 && input[4] == 5 {
		return 15
	}
	return 0
}
func largestraight(input []int) int {
	sort.Ints(input)
	if input[0] == 2 && input[1] == 3 && input[2] == 4 && input[3] == 5 && input[4] == 6 {
		return 20
	}
	return 0
}

func fullhouse(input []int) int {
	sort.Ints(input)
	if input[0] == input[1] && input[0] == input[2] && input[0] != input[3] && input[3] == input[4] {
		return 3*input[0] + 2*input[3]
	}
	if input[0] == input[1] && input[0] != input[2] && input[2] == input[3] && input[3] == input[4] {
		return 2*input[0] + 3*input[3]
	}
	return 0
}

func chance(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum
}
