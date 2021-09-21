package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	// Get the line as array
	words := strings.Fields(sc.Text())

	// Finds first and second word, and changes lower / upper case
	first := strings.ToLower(words[0])
	second := strings.Title(words[1])

	// Reverses first and second word
	words[0] = second
	words[1] = first

	// Replace in the last word '?' with '!'
	words[len(words)-1] = strings.ReplaceAll(words[len(words)-1], "?", "!")

	fmt.Println(strings.Join(words, " "))
}
