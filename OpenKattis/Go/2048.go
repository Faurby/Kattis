package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var board [4][4]int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sc.Scan()
			board[i][j], _ = strconv.Atoi(sc.Text())
		}
	}

	sc.Scan()
	direction, _ := strconv.Atoi(sc.Text())

	// fmt.Printf("Before\n")
	// display()
	switch direction {
	case 0:
		goLeft()
	case 1:
		goUp()
	case 2:
		goRight()
	case 3:
		goDown()
	}

	// fmt.Printf("After\n")
	display()
}

func goLeft() {
	for i := 0; i < 4; i++ {
		counter := 0

		// Move everything left
		for j := 0; j < 4; j++ {
			current := &board[i][j]

			if *current == 0 {
				counter++
			} else if *current != 0 && counter != 0 {
				board[i][j-counter] = *current
				*current = 0
			}
		}

		// Identify identical neighbors in one row,
		// and double the left one, set right one zero
		for j := 0; j < 3; j++ {
			current := &board[i][j]
			next := &board[i][j+1]
			if *current == *next && *current != 0 {
				*current = 2 * *current
				*next = 0
				// If neighbors have been merged, move everything 1 to the left
				for k := j + 1; k < 3; k++ {
					board[i][k] = board[i][k+1]
				}
				// Since we cant do board[i][k+1] on the last index, do this
				board[i][3] = 0
			}
		}
	}
}

func goUp() {
	for i := 0; i < 4; i++ {
		counter := 0

		// Move everything left
		for j := 0; j < 4; j++ {
			current := &board[j][i]

			if *current == 0 {
				counter++
			} else if *current != 0 && counter != 0 {
				board[j-counter][i] = *current
				*current = 0
			}
		}

		// Identify identical neighbors in one row,
		// and double the left one, set right one zero
		for j := 0; j < 3; j++ {
			current := &board[j][i]
			next := &board[j+1][i]
			if *current == *next && *current != 0 {
				*current = 2 * *current
				*next = 0
				// If neighbors have been merged, move everything 1 to the left
				for k := j + 1; k < 3; k++ {
					board[k][i] = board[k+1][i]
				}
				// Since we cant do board[i][k+1] on the last index, do this
				board[3][i] = 0
			}
		}
	}
}

func goRight() {
	for i := 0; i < 4; i++ {
		counter := 0

		// Move everything left
		for j := 3; j > -1; j-- {
			current := &board[i][j]

			if *current == 0 {
				counter++
			} else if *current != 0 && counter != 0 {
				board[i][j+counter] = *current
				*current = 0
			}
		}

		// Identify identical neighbors in one row,
		// and double the left one, set right one zero
		for j := 3; j > 0; j-- {
			current := &board[i][j]
			next := &board[i][j-1]
			if *current == *next && *current != 0 {
				*current = 2 * *current
				*next = 0
				// If neighbors have been merged, move everything 1 to the left
				for k := j - 1; k > 0; k-- {
					board[i][k] = board[i][k-1]
				}
				// Since we cant do board[i][k+1] on the last index, do this
				board[i][0] = 0
			}
		}
	}
}

func goDown() {
	for i := 0; i < 4; i++ {
		counter := 0

		// Move everything left
		for j := 3; j > -1; j-- {
			current := &board[j][i]

			if *current == 0 {
				counter++
			} else if *current != 0 && counter != 0 {
				board[j+counter][i] = *current
				*current = 0
			}
		}

		// Identify identical neighbors in one row,
		// and double the left one, set right one zero
		for j := 3; j > 0; j-- {
			current := &board[j][i]
			next := &board[j-1][i]
			if *current == *next && *current != 0 {
				*current = 2 * *current
				*next = 0
				// If neighbors have been merged, move everything 1 to the left
				for k := j - 1; k > 0; k-- {
					board[k][i] = board[k-1][i]
				}
				// Since we cant do board[i][k+1] on the last index, do this
				board[0][i] = 0
			}
		}
	}
}

func display() {
	for i := 0; i < 4; i++ {

		for j := 0; j < 4; j++ {
			fmt.Printf("%v ", board[i][j])
		}
		fmt.Println()
	}
}
