package day1

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

	func getInput() *[2][]int {
		var (
			id_list_str [2][]string
			id_list [2][]int
			data []byte
			txt_data string
		)
		data, err := os.ReadFile("./day1/input.txt")
		if err != nil {
			log.Printf("Could not open file.\n error: %s", err.Error())
			return nil
		}

		txt_data = string(data)
		
		lines := strings.Split(txt_data, "\n")
		for _, line := range(lines) {
			line_ids := strings.Split(line, " ")

			if len(line_ids) != 2 {
				for i := len(line_ids)-1; i>=0; i-- {
					if line_ids[i] == "" { 
						line_ids = append(line_ids[:i], line_ids[i+1:]...)
					}
				}
			}

			id_list_str[0] = append(id_list_str[0], line_ids[0])
			id_list_str[1] = append(id_list_str[1], line_ids[1])
		}

		for i, s := range(id_list_str) {
			for _, str := range(s) {
				id, err := strconv.Atoi(str)
				if err != nil {
					log.Println("error parsing num:", str)
					return nil
				}

				id_list[i] = append(id_list[i], id)
			}
		}

		return &id_list
	}

	func Solve() {
		list := *getInput()

		for _, s := range(list) {
			slices.SortFunc(s, func(a int, b int) int {return a-b})
		}

		distance := 0

		for i := range(list[0]) {
			first := list[0][i]
			second := list[1][i]

			if first > second {
				distance += first - second
			} else {
				distance += second - first
			}
		}

		fmt.Printf("Total distance is: %d\n", distance)
	}