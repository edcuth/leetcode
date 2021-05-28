//116. Populating Next Right Pointers in Each Node
//https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	next := root

	for next.Left != nil {
		current := next
		for current != nil {
			current.Left.Next = current.Right
			if current.Next != nil {
				current.Right.Next = current.Next.Left
			}
			current = current.Next
		}
		next = next.Left
	}
	return root
}