//104. Maximum Depth of Binary Tree
//https://leetcode.com/problems/maximum-depth-of-binary-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count1 := 1
	count2 := 1
	count1 += maxDepth(root.Left)
	count2 += maxDepth(root.Right)
	if count1 > count2 {
		return count1
	}
	return count2
}

