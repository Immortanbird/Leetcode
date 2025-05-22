package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	return traverse(root)
}

func traverse(root *TreeNode) []int {
	seq := []int{}
	if root.Left != nil {
		seq = append(seq, traverse(root.Left)...)
	}
	seq = append(seq, root.Val)
	if root.Right != nil {
		seq = append(seq, traverse(root.Right)...)
	}

	return seq
}
