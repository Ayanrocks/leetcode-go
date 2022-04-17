package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	result := findRelativeRanks([]int{5, 4, 3, 2, 1})
	fmt.Println("Result: ", result)
}

// Implement priority queue

type Item struct {
	Priority int

	Index int //. index of the original array
}

type PQ []*Item

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[j].Priority < pq[i].Priority // descending order
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]

	*pq = (*pq)[:len(*pq)-1]
	return item
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func findRelativeRanks(score []int) []string {
	// create a priority queue and insert element
	pq := make(PQ, len(score))
	res := make([]string, len(score))
	idx := 0
	for i, v := range score {
		pq[idx] = &Item{
			Priority: v,
			Index:    i,
		}
		idx++
	}

	heap.Init(&pq)
	heap.Push()

	count := 1

	for _ = range score {
		if count == 1 {
			top := heap.Pop(&pq).(*Item)
			res[top.Index] = "Gold Medal"
		} else if count == 2 {
			top := heap.Pop(&pq).(*Item)
			res[top.Index] = "Silver Medal"
		} else if count == 3 {
			top := heap.Pop(&pq).(*Item)
			res[top.Index] = "Bronze Medal"
		} else {
			top := heap.Pop(&pq).(*Item)
			res[top.Index] = strconv.Itoa(top.Priority)
		}

		count++
	}

	return res
}
