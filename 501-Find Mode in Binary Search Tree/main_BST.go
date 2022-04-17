package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	prev := -1 << 31
	count := 0
	max := 0
	result := []int{}
	findBSTMode(root, &prev, &max, &count, &result)
	return result
}

func findBSTMode(root *TreeNode, prev, max, count *int, result *[]int) {
	if root == nil {
		return
	}

	findBSTMode(root.Left, prev, max, count, result)

	if root.Val == *prev {
		*count++
	} else {
		*count = 1
	}

	if *max < *count {
		*result = []int{root.Val}
		*max = *count
	} else if *max == *count {
		*result = append(*result, root.Val)
	}

	*prev = root.Val

	findBSTMode(root.Right, prev, max, count, result)
}
