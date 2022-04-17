package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	result := findRelativeRanks([]int{10, 3, 8, 9, 4})
	fmt.Println("Result: ", result)
}

func findRelativeRanks(score []int) []string {
	// create a map
	m := map[int]int{}
	res := make([]string, len(score))

	for i, v := range score {
		m[v] = i
	}

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	fmt.Println(score)

	for i := 0; i < len(score); i++ {
		if i == 0 {
			res[m[score[i]]] = "Gold Medal"
		} else if i == 1 {
			res[m[score[i]]] = "Silver Medal"
		} else if i == 2 {
			res[m[score[i]]] = "Bronze Medal"
		} else {
			res[m[score[i]]] = strconv.Itoa(i + 1)
		}
	}

	return res
}
