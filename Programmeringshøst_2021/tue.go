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
	var curpos, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	var wishpos, _ = strconv.Atoi(sc.Text())

	var dist int

	if curpos < wishpos {
		if wishpos-curpos >= 180 {
			dist = -1 * (360 - (wishpos - curpos))
		} else {
			dist = wishpos - curpos
		}
	} else {
		if curpos-wishpos >= 180 {
			dist = 360 - (curpos - wishpos)
		} else {
			dist = -1 * (curpos - wishpos)
		}
	}
	fmt.Println(dist)
}
