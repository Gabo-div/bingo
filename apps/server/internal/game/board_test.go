package game

import (
	"math/rand"
	"testing"
	"time"
)

func TestGenerateRandomBoard(t *testing.T) {
	boardA, _ := GenerateBoard(0)
	boardB, _ := GenerateBoard(0)
	equal := true

	for j := range 5 {
		for i := range 5 {
			if boardA[j][i] != boardB[j][i] {
				equal = false
			}
		}
	}

	if equal {
		t.Errorf("does not generate a different board from random seed")
	}
}

func TestGenerateSeedBoard(t *testing.T) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	seed := r.Int63()

	boardA, _ := GenerateBoard(seed)
	boardB, _ := GenerateBoard(seed)
	equal := true

	for j := range 5 {
		for i := range 5 {
			if boardA[j][i] != boardB[j][i] {
				equal = false
			}
		}
	}

	if !equal {
		t.Errorf("does not generate the same board for seed: %d", seed)
	}
}

func TestCheckHorizontalWin(t *testing.T) {
	board := [][]int{
		{2, 16, 34, 47, 61},
		{7, 20, 33, 41, 70},
		{14, 25, 0, 56, 64},
		{8, 22, 36, 59, 65},
		{1, 26, 44, 48, 68},
	}

	if !CheckHorizontalWin(board, []int{8, 22, 36, 59, 65}) {
		t.Errorf("does not check horizontal win condition")
	}

	if !CheckHorizontalWin(board, []int{14, 25, 56, 64}) {
		t.Errorf("does not check horizontal win condition on middle")
	}

	if CheckHorizontalWin(board, []int{14, 25, 56}) {
		t.Errorf("does check horizontal win condition when it should not")
	}
}

func TestCheckVerticalWin(t *testing.T) {
	board := [][]int{
		{2, 16, 34, 47, 61},
		{7, 20, 33, 41, 70},
		{14, 25, 0, 56, 64},
		{8, 22, 36, 59, 65},
		{1, 26, 44, 48, 68},
	}

	if !CheckVerticalWin(board, []int{34, 33, 36, 44}) {
		t.Errorf("does not check vertical win condition")
	}

	if !CheckVerticalWin(board, []int{47, 41, 56, 59, 48}) {
		t.Errorf("does not check vertical win condition on middle")
	}

	if CheckVerticalWin(board, []int{47, 41, 56, 59}) {
		t.Errorf("does check vertical win condition when it should not")
	}
}

func TestLeftDiagonalWin(t *testing.T) {
	board := [][]int{
		{2, 16, 34, 47, 61},
		{7, 20, 33, 41, 70},
		{14, 25, 0, 56, 64},
		{8, 22, 36, 59, 65},
		{1, 26, 44, 48, 68},
	}

	if !CheckLeftDiagonalWin(board, []int{2, 20, 59, 68}) {
		t.Errorf("does not check left diagonal win condition")
	}

	if CheckLeftDiagonalWin(board, []int{2, 20, 59}) {
		t.Errorf("does check left diagonal win condition when it should not")
	}
}

func TestRightDiagonalWin(t *testing.T) {
	board := [][]int{
		{2, 16, 34, 47, 61},
		{7, 20, 33, 41, 70},
		{14, 25, 0, 56, 64},
		{8, 22, 36, 59, 65},
		{1, 26, 44, 48, 68},
	}

	if !CheckRightDiagonalWin(board, []int{61, 41, 22, 1}) {
		t.Errorf("does not check right diagonal win condition")
	}

	if CheckRightDiagonalWin(board, []int{61, 41, 22}) {
		t.Errorf("does check right diagonal win condition when it should not")
	}
}

func BenchmarkGenerateBoard(t *testing.B) {
	for t.Loop() {
		GenerateBoard(0)
	}
}

func BenchmarkCheckWin(t *testing.B) {
	board := [][]int{
		{2, 16, 34, 47, 61},
		{7, 20, 33, 41, 70},
		{14, 25, 0, 56, 64},
		{8, 22, 36, 59, 65},
		{1, 26, 44, 48, 68},
	}

	for t.Loop() {
		CheckWin(board, []int{100, 101, 102, 103, 104, 105, 106, 107})
	}
}
