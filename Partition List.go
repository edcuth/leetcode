//86. Partition List
//https://leetcode.com/problems/partition-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	headLess := &ListNode{Val: 0}
	nextLess := headLess
	headMore := &ListNode{Val: 0}
	nextMore := headMore
	current := head
	for current.Next != nil {
		if current.Val < x {
			nextLess.Next = current
			nextLess = nextLess.Next
			current = current.Next
		} else {
			nextMore.Next = current
			nextMore = nextMore.Next
			current = current.Next
		}
	}
	if current.Val < x {
		nextLess.Next = current
		nextLess = nextLess.Next
		nextMore.Next = nil
	} else {
		nextMore.Next = current
		nextMore = nextMore.Next
		nextMore.Next = nil
	}
	nextLess.Next = headMore.Next
	return headLess.Next
}