//226. Invert Binary Tree
//https://leetcode.com/problems/invert-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}

func invert(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	invert(node.Left)
	invert(node.Right)
}