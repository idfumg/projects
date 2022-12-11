package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	PLAYER   = "x"
	COMPUTER = "o"
)

func isValid(grid [][]string, i int, j int) bool {
	return (i < len(grid) &&
		j < len(grid) &&
		i >= 0 &&
		j >= 0 &&
		grid[i][j] != "x" &&
		grid[i][j] != "o")
}

func readRowAndCol(grid [][]string) (int, int) {
	for {
		row := 0
		fmt.Print("Row: ")
		fmt.Scanln(&row)

		col := 0
		fmt.Print("Col: ")
		fmt.Scanln(&col)

		if !isValid(grid, row, col) {
			fmt.Println("Wrong input: ", row, col)
			continue
		}

		return row, col
	}
}

func wonByColumn(grid [][]string) bool {
	for j := 0; j < len(grid); j += 1 {
		cnt1 := 0
		cnt2 := 0
		for i := 0; i < len(grid); i += 1 {
			if grid[i][j] == "x" {
				cnt1 += 1
			} else if grid[i][j] == "o" {
				cnt2 += 1
			}
		}
		if cnt1 == len(grid) || cnt2 == len(grid) {
			return true
		}
	}
	return false
}

func wonByRow(grid [][]string) bool {
	for i := 0; i < len(grid); i += 1 {
		cnt1 := 0
		cnt2 := 0
		for j := 0; j < len(grid); j += 1 {
			if grid[i][j] == "x" {
				cnt1 += 1
			} else if grid[i][j] == "o" {
				cnt2 += 1
			}
		}
		if cnt1 == len(grid) || cnt2 == len(grid) {
			return true
		}
	}
	return false
}

func wonByLeftDiagonal(grid [][]string) bool {
	cnt1 := 0
	cnt2 := 0
	for i := 0; i < len(grid); i += 1 {
		if grid[i][i] == "x" {
			cnt1 += 1
		} else if grid[i][i] == "o" {
			cnt2 += 1
		}
	}
	return cnt1 == len(grid) || cnt2 == len(grid)
}

func wonByRightDiagonal(grid [][]string) bool {
	cnt1 := 0
	cnt2 := 0
	for j := 0; j < len(grid); j += 1 {
		if grid[len(grid)-j-1][j] == "x" {
			cnt1 += 1
		} else if grid[len(grid)-j-1][j] == "o" {
			cnt2 += 1
		}
	}
	return cnt1 == len(grid) || cnt2 == len(grid)
}

func isThereAWin(grid [][]string) bool {
	return (wonByColumn(grid) ||
		wonByRow(grid) ||
		wonByLeftDiagonal(grid) ||
		wonByRightDiagonal(grid))
}

func generateRowAndCol(grid [][]string) (int, int) {
	rand.Seed(time.Now().UnixNano())
	for {
		i := rand.Intn(len(grid))
		j := rand.Intn(len(grid))
		if isValid(grid, i, j) {
			return i, j
		}
	}
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i += 1 {
		fmt.Println(grid[i])
	}
}

func printAndCheckGrid(grid [][]string, row int, col int, p string) bool {
	grid[row][col] = p
	printGrid(grid)
	return isThereAWin(grid)
}

func start() {
	grid := [][]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}

	printGrid(grid)

	for {
		row, col := readRowAndCol(grid)
		if printAndCheckGrid(grid, row, col, PLAYER) {
			fmt.Println("Player wins")
			return
		}

		row, col = generateRowAndCol(grid)
		if printAndCheckGrid(grid, row, col, COMPUTER) {
			fmt.Println("Computer wins")
			return
		}
	}
}

func main() {
	// start()
}
