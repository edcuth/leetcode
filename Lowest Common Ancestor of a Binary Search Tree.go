//235. Lowest Common Ancestor of a Binary Search Tree
//https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for true {
		if q.Val < root.Val && p.Val < root.Val {
			root = root.Left
		} else if q.Val > root.Val && p.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}