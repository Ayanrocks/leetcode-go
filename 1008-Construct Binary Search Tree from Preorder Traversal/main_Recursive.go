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

func bstFromPreorder(preorder []int) *TreeNode {
	i := 0
	return reconstruct(preorder, &i, 1<<32)
}

func reconstruct(preOrder []int, i *int, bound int) *TreeNode {
	if *i == len(preOrder) || preOrder[*i] > bound {
		return nil
	}

	root := &TreeNode{Val: preOrder[*i]}
	*i++
	root.Left = reconstruct(preOrder, i, root.Val)
	root.Right = reconstruct(preOrder, i, bound)
	return root
}
