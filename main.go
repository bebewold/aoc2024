package main

import (
	"aoc2024/day1"
	"aoc2024/day2"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	println("args:", args)
	choice := strings.Join(args, " ")
	again := true
	for again {
		switch choice {
		case "day1" : 
			var part string
			print("part1 or part2 ?")
			fmt.Scanf("%s\n", &part)
			choice += " "+part

		case "day1 part1":
			day1.Solve()
			again = false

		case "day1 part2":
			day1.Solve_part2()
			again = false

		case "day2 part1":
			day2.Solve_part1()
			again = false

		default: 
			println("Wrong choice...\n exiting")
			println(choice)
			os.Exit(3)
			again = false
		}
	}
}