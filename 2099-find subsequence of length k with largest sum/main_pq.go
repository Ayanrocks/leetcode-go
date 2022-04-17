package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	result := maxSubsequence([]int{2, 1, 3, 3}, 2)
	fmt.Println(result)
}

type Item struct {
	Val, Priority int
}

type PQ []Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func maxSubsequence(nums []int, k int) []int {
	pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < k; i++ {
		item := Item{
			Val:      i,
			Priority: nums[i],
		}

		heap.Push(pq, item)
	}

	for i := k; i < len(nums); i++ {
		if (*pq)[0].Priority < nums[i] {
			heap.Pop(pq)
			heap.Push(pq, Item{Priority: nums[i], Val: i})
		}
	}

	sort.SliceStable(*pq, func(i, j int) bool {
		return (*pq)[i].Val < (*pq)[j].Val
	})

	ans := []int{}
	for _, v := range *pq {
		ans = append(ans, v.Priority)
	}

	return ans
}
