package day1

import "fmt"

func Solve_part2() {
	var(
		similarity = 0
	)
	list := getInput()

	left_list := list[0]
	right_list := list[1]

	for _, left_id := range left_list {
		for _, right_id := range right_list {
			if left_id == right_id {
				similarity += left_id
			}
		}
	}

	fmt.Printf("Similarity=%d\n", similarity)
}