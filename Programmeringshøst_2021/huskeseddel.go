package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()

	crypted, _ := strconv.Atoi(sc.Text())
	crypted = (crypted + 1) % 10000
	resultString := fmt.Sprintf("%04d", crypted)

	resultStringArr := strings.Split(resultString, "")
	reversedResultArr := reverse(resultStringArr)

	fmt.Println(strings.Join(reversedResultArr, ""))
}

func reverse(numbers []string) []string {
	newNumbers := make([]string, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}
