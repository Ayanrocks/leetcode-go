package main

import "container/heap"

// Implement priority queue
type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq PQ) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

type KthLargest struct {
	KthVal int
	PQ     *PQ
}

func Constructor(k int, nums []int) KthLargest {
	var h KthLargest

	if len(nums) < k {
		h = KthLargest{PQ: (*PQ)(&nums), KthVal: k}
		heap.Init(h.PQ)
	} else {
		kNums := nums[:k]
		h = KthLargest{PQ: (*PQ)(&kNums), KthVal: k}
		heap.Init(h.PQ)
		for i := k; i < len(nums); i++ {
			h.Add(nums[i])
		}
	}

	return h
}

func (this *KthLargest) Add(val int) int {
	// add value to PQ

	if this.PQ.Len() < this.KthVal {
		heap.Push(this.PQ, val)
	} else if val > (*this.PQ)[0] {
		(*this.PQ)[0] = val
		heap.Fix(this.PQ, 0)
	}

	return (*this.PQ)[0]
}
