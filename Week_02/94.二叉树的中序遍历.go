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
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
// 递归解法
var tree []int
func inorderTraversal(root *TreeNode) []int {
	dfs(root)
	return tree
}
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	tree = append(tree, root.Val)
	dfs(root.Right)
	return
}

// 遍历解法
var stack []*TreeNode
func inorderTraversal(root *TreeNode) []int {
	tree = []int{}
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}else {
			lastIndex := len(stack) - 1
			tmp := stack[lastIndex]
			tree = append(tree, tmp.Val)
			stack = stack[:lastIndex]
			root = tmp.Right
		}
	}
	return tree
}
