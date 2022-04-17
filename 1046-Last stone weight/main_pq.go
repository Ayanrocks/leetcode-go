package main

import (
	"container/heap"
	"fmt"
)

func main() {
	res := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	fmt.Println("Res: ", res)
}

type PQ []int

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	elem := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return elem
}

func lastStoneWeight(stones []int) int {
	pq := (*PQ)(&stones)

	heap.Init(pq)

	for pq.Len() > 1 {
		newNode := heap.Pop(pq).(int) - heap.Pop(pq).(int)

		if newNode > 0 {
			heap.Push(pq, newNode)
		}
	}

	if pq.Len() == 1 {
		return heap.Pop(pq).(int)
	}

	return 0
}