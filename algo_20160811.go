package main

import (
	"fmt"
	"strconv"
)

func main() {

	board := [][]string{
		{
			"*", ".", ".", ".",
		},
		{
			".", ".", ".", ".",
		},
		{
			".", "*", ".", ".",
		},
		{
			".", ".", ".", ".",
		},
	}
	fmt.Println(find(board))

	board = [][]string{
		{"*", "*", ".", ".", "."},
		{".", ".", ".", ".", "."},
		{".", "*", ".", ".", "."},
	}
	fmt.Println(find(board))
}

func find(b [][]string) [][]string {
	answer := make([][]string, len(b), (cap(b)+1)*2)
	copy(answer, b)
	for r, row := range b {
		for c, cell := range row {
			if cell == "." {
				mineCount := 0
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						if i == 0 && j == 0 {
							continue
						}
						checkRow := r + i
						checkCell := c + j
						if (checkRow >= 0 && checkRow < len(b)) && (checkCell >= 0 && checkCell < len(row)) {
							if b[checkRow][checkCell] == "*" {
								mineCount++
							}
						}
					}
				}
				answer[r][c] = strconv.Itoa(mineCount)
			}
		}
	}
	return answer
}
