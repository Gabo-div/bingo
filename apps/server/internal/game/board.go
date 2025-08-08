package game

import (
	"math/rand"
	"slices"
	"time"
)

func GenerateBoard(seed int64) ([][]int, int64) {
	finalSeed := seed

	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	r := rand.New(rand.NewSource(finalSeed))

	columnsRanges := map[string]struct {
		start int
		end   int
	}{
		"B": {1, 15},
		"I": {16, 30},
		"N": {31, 45},
		"G": {46, 60},
		"O": {61, 75},
	}

	columnsLetters := []string{"B", "I", "N", "G", "O"}

	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	for i, letra := range columnsLetters {
		rangoInfo := columnsRanges[letra]

		numbers := make([]int, 0, rangoInfo.end-rangoInfo.start+1)
		for j := rangoInfo.start; j <= rangoInfo.end; j++ {
			numbers = append(numbers, j)
		}

		r.Shuffle(len(numbers), func(k, l int) {
			numbers[k], numbers[l] = numbers[l], numbers[k]
		})

		board[i] = numbers[:5]
	}

	finalBoard := make([][]int, 5)
	for i := range finalBoard {
		finalBoard[i] = make([]int, 5)
		for j := range finalBoard[0] {
			finalBoard[i][j] = board[j][i]
		}
	}

	finalBoard[2][2] = 0

	return finalBoard, finalSeed
}

func CheckHorizontalWin(board [][]int, numbers []int) bool {
	for j := range 5 {

		win := true

		for i := range 5 {

			n := board[j][i]

			if n != 0 && !slices.Contains(numbers, n) {
				win = false
			}
		}

		if win {
			return true
		}
	}

	return false
}

func CheckVerticalWin(board [][]int, numbers []int) bool {
	for i := range 5 {
		win := true

		for j := range 5 {
			n := board[j][i]

			if n != 0 && !slices.Contains(numbers, n) {
				win = false
			}
		}

		if win {
			return true
		}
	}

	return false
}

func CheckLeftDiagonalWin(board [][]int, numbers []int) bool {
	for i := range 5 {
		n := board[i][i]

		if n != 0 && !slices.Contains(numbers, n) {
			return false
		}
	}

	return true
}

func CheckRightDiagonalWin(board [][]int, numbers []int) bool {
	for i := range 5 {
		n := board[i][len(board)-1-i]

		if n != 0 && !slices.Contains(numbers, n) {
			return false
		}
	}

	return true
}

func CheckWin(board [][]int, numbers []int) bool {
	return CheckHorizontalWin(board, numbers) ||
		CheckVerticalWin(board, numbers) ||
		CheckLeftDiagonalWin(board, numbers) ||
		CheckRightDiagonalWin(board, numbers)
}
