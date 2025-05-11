package day2

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type level int
type report []level

func Solve_part1() {
	reports := getInput()
	count := 0

	for _, rep := range reports {
		print("report len:", len([]level(rep)))
		if rep.isSafe() {
			count++
		}
	}
	println("Count:", count)
}

func (rep report) isSafe() (safe bool) {
	increasing := rep[0] < rep[1]

	for i := range rep[:len(rep)-1] {
		println("doing line", i, "of", len(rep))
		if diff := math.Abs(float64(rep[i] - rep[i+1])); diff > 3 {
			safe = false
			return
		}

		if increasing {
			if rep[i] >= rep[i+1] {
				return
			}

		} else if rep[i] <= rep[i+1] {
			return
		}

	}
	safe = true
	return 
}

func getInput() []report {
	var reports []report = []report{}

	data, err := os.ReadFile("./day2/input.txt")
	if err != nil {
		log.Printf("Could not open file.\n error: %s", err.Error())
		return nil
	}

	txt_data := string(data)
	
	lines := strings.Split(txt_data, "\n")
	for _, line := range(lines) {
		var rep report = report{}
		words := strings.Split(line, " ")

		if len(words) != 1 {
			for i := len(words)-1; i>=0; i-- {
				if words[i] == "" { 
					words = append(words[i:], words[:i+1]... )
				} else {
					num, err := strconv.Atoi(words[i])
					if err != nil {
						log.Panicln("Could not parse day2 data to int")
						return nil
					}

					rep = append(report{level(num)}, rep...)
				}
			}
		}

		reports = append(reports, rep)
	}

	return reports
}