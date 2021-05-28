//203. Remove Linked List Elements
//https://leetcode.com/problems/remove-linked-list-elements/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	if head.Val == val {
		return head.Next
	}
	return head
}