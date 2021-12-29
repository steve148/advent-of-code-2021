package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	numbers [][]int
	marked  [5][5]bool
}

func (b *BingoBoard) updateBoard(x int) {
	for i := 0; i < len(b.numbers); i++ {
		for j := 0; j < len(b.numbers[i]); j++ {
			if b.numbers[i][j] == x {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *BingoBoard) isBingo() bool {
	// Do any rows have a full columns
	for i := 0; i < len(b.marked); i++ {
		if isBingo(b.marked[i][:]) {
			return true
		}
	}
	// Do any columns have a match
	for i := 0; i < len(b.marked); i++ {
		var column []bool
		for _, row := range b.marked {
			column = append(column, row[i])
		}
		if isBingo(column) {
			return true
		}
	}

	return false
}

func isBingo(b []bool) bool {
	for _, v := range b {
		if !v {
			return false
		}
	}
	return true
}

func (b *BingoBoard) sumUnmarked() int {
	sum := 0
	for i := 0; i < len(b.marked); i++ {
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}

func playGame(numbers []int, boards []*BingoBoard) int {
	for _, number := range numbers {
		for _, board := range boards {
			board.updateBoard(number)
			if board.isBingo() {
				return board.sumUnmarked() * number
			}

		}
	}
	log.Fatal("No winner found")
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var chosen_numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		chosen_numbers = append(chosen_numbers, i)
	}

	var boards []*BingoBoard

	var numbers [][]int
	for _, line := range lines[1:] {
		trimmed := strings.TrimSpace(line)
		row_strings := strings.Fields(trimmed)
		var row_numbers []int
		for _, s := range row_strings {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			row_numbers = append(row_numbers, i)
		}

		if line != "" {
			numbers = append(numbers, row_numbers)
		}

		if len(numbers) == 5 {
			var marked [5][5]bool
			board := BingoBoard{numbers: numbers, marked: marked}
			boards = append(boards, &board)
			numbers = nil
		}
	}

	score := playGame(chosen_numbers, boards)

	fmt.Println(score)
}
