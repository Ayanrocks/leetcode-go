package main

import (
	"container/heap"
	"fmt"
)

func main() {
	result := findKthLargest([]int{2, 1, 3, 3}, 2)
	fmt.Println(result)
}

type PQ []int

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[j] < pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func findKthLargest(nums []int, k int) int {
	pq := (*PQ)(&nums)

	heap.Init(pq)
	ans := 0
	for i := 0; i < k; i++ {
		ans = heap.Pop(pq).(int)
	}

	return ans
}
