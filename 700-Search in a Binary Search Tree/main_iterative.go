package main

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

func increasingBST(root *TreeNode) *TreeNode {
	s := []*TreeNode{}
	curr := root
	head := &TreeNode{Val: -1}
	p := head
	for len(s) != 0 || curr != nil {
		// traversing left and pushin curr to stack
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}

		// popping from stack
		curr = s[len(s)-1]
		s = s[:len(s)-1]

		curr.Left = nil
		p.Right = curr
		p = curr
		curr = curr.Right
	}

	return head.Right
}
