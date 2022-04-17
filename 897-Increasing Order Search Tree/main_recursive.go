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
	return reArrange(root, nil)
}

func reArrange(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}

	res := reArrange(root.Left, root)
	root.Left = nil
	root.Right = reArrange(root.Right, tail)
	return res
}
