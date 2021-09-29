package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var coordinates []float64

	//create slice of coordinates
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.ParseFloat(sc.Text(), 32)

		coordinates = append(coordinates, x)

		sc.Scan()
		y, _ := strconv.ParseFloat(sc.Text(), 32)
		coordinates = append(coordinates, y)
	}
	//insert first x/y last in slice
	firstX := coordinates[0]
	firstY := coordinates[1]
	coordinates = append(coordinates, firstX, firstY)

	//using area of a polygon, coordinate geometry
	divident := 0.0
	for i := 0; i < len(coordinates)-2; i += 2 {
		x1 := coordinates[i]
		y1 := coordinates[i+1]
		x2 := coordinates[i+2]
		y2 := coordinates[i+3]

		divident += (x1 * y2) - (y1 * x2)
	}

	fmt.Printf("%.1f", math.Abs(divident/2))
}
