package main

import "fmt"

func main() {
	result := findMode(&TreeNode{Val: 0})
	fmt.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m = make(map[int]int, 0)

func findMode(root *TreeNode) []int {
	traverse(root)
	max := 0
	ans := []int{}

	for k, v := range m {
		if v == max {
			ans = append(ans, k)
		} else if v > max {
			ans = []int{k}
			max = v
		}
	}

	return ans
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}

	m[root.Val]++
	traverse(root.Left)
	traverse(root.Right)
}
