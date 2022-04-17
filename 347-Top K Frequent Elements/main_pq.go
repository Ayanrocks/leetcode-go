package main

import (
	"container/heap"
	"fmt"
)

func main() {
	result := topKFrequent([]int{2, 1, 3, 3}, 2)
	fmt.Println(result)
}

type IntHeap struct {
	Val, Count int
}

type PQ []IntHeap

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[j].Count < pq[i].Count
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(IntHeap))
}
func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
	}

	pq := &PQ{}
	heap.Init(pq)

	for k, v := range m {
		heap.Push(pq, IntHeap{Val: k, Count: v})
	}

	ans := []int{}

	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(pq).(IntHeap).Val)
	}
	

	return ans

}
