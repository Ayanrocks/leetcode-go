package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root
	val1, val2 := p.Val, q.Val
	for {
		switch {
		case curr.Val < val1 && curr.Val < val2:
			curr = curr.Right
		case curr.Val > val1 && curr.Val > val2:
			curr = curr.Left
		default:
			return curr
		}
	}
}
