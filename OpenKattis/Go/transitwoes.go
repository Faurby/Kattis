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

	// Leave home at time s
	currentTime, _ := strconv.Atoi(sc.Text())

	sc.Scan()

	// Class start at time t. if s is below t, then print yes. Otherwise print no. (if he makes it to class)
	classStart, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// walkTo[0] is walk to first bus, and walkTo[i] is to ith bus. walkTo[len(walkTo)-1] = walk from las bus to class
	var walkTo []int
	for i := 0; i < n+1; i++ {
		sc.Scan()
		walkTime, _ := strconv.Atoi(sc.Text())
		walkTo = append(walkTo, walkTime)
	}

	var busRide []int
	for i := 0; i < n; i++ {
		sc.Scan()
		busTime, _ := strconv.Atoi(sc.Text())
		busRide = append(busRide, busTime)
	}

	var arrivals []int
	for i := 0; i < n; i++ {
		sc.Scan()
		arrival, _ := strconv.Atoi(sc.Text())
		arrivals = append(arrivals, arrival)
	}

	for i := 0; i < n+1; i++ {
		// First walk to first bus
		currentTime += walkTo[i]

		if i == len(walkTo)-1 {
			if currentTime < classStart {
				fmt.Println("yes")
				return
			} else {
				fmt.Println("no")
				return
			}
		}
		// Wait for first bus to arrive
		currentTime += currentTime % arrivals[i]
		// Then ride first bus
		currentTime += busRide[i]
	}

}
