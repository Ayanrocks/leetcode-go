package main

import (
	"container/heap"
	"fmt"
)

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

	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				count++
			}

			if mat[i][j] == 0 {
				break
			}
		}

		pq[i] = Item{
			Count: count,
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
