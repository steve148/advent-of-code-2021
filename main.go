package main

import (
	"os"

	"github.com/steve148/advent-of-code-2021/day5"
	day5 "github.com/steve148/advent-of-code-2021/day5"
)

func main() {
	args := os.Args[1:]
	day := args[0]
	part := args[1]

	if day == "1" && part == "1" {
		day5.Part1()
	} else if day == "1" && part == "2" {
		day5.Part2()
	}
}
