package main

import (
	"container/heap"
	"fmt"
)

// With priority queue and binary search
func main() {
	res := kWeakestRows([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}, 3)
	fmt.Println(res)

	res = kWeakestRows([][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}, 2)
	fmt.Println(res)

	// [2 0 3]
	// [0 2]
}

type Item struct {
	Count int
	Row   int
}

type PQ []Item

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool {
	if pq[i].Count == pq[j].Count {
		return pq[i].Count+pq[i].Row < pq[j].Count+pq[j].Row
	}
	return pq[i].Count < pq[j].Count
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PQ) Pop() interface{} {
	elem := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return elem
}

func kWeakestRows(mat [][]int, k int) []int {
	pq := make(PQ, len(mat))

	// binary search
	for i := 0; i < len(mat); i++ {
		lo := 0
		hi := len(mat[i])
		for lo < hi {
			mid := lo + (hi-lo)/2

			if mat[i][mid] == 1 {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

		pq[i] = Item{
			Count: lo,
			Row:   i,
		}
	}

	heap.Init(&pq)

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&pq).(Item).Row)
	}

	return res
}
