package main

import (
	"os"

	day4 "github.com/steve148/advent-of-code-2021/day4"
	day5 "github.com/steve148/advent-of-code-2021/day5"
	day6 "github.com/steve148/advent-of-code-2021/day6"
	"github.com/steve148/advent-of-code-2021/day7"
)

func main() {
	args := os.Args[1:]
	day := args[0]
	part := args[1]

	if day == "4" && part == "1" {
		day4.Part1()
	} else if day == "4" && part == "2" {
		day4.Part2()
	} else if day == "5" && part == "1" {
		day5.Part1()
	} else if day == "5" && part == "2" {
		day5.Part2()
	} else if day == "6" && part == "1" {
		day6.Part1()
	} else if day == "6" && part == "2" {
		day6.Part2()
	} else if day == "7" && part == "1" {
		day7.Part1()
	} else if day == "7" && part == "2" {
		day7.Part2()
	}
}
