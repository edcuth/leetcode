//21. Merge Two Sorted Lists
//https://leetcode.com/problems/merge-two-sorted-lists/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	current := result
	i := l1
	j := l2
	for i != nil && j != nil {
		if i.Val > j.Val {
			current.Next = &ListNode{Val: j.Val}
			current = current.Next
			j = j.Next
		} else {
			current.Next = &ListNode{Val: i.Val}
			current = current.Next
			i = i.Next
		}
	}
	for i != nil {
		current.Next = &ListNode{Val: i.Val}
		current = current.Next
		i = i.Next
	}
	for j != nil {
		current.Next = &ListNode{Val: j.Val}
		current = current.Next
		j = j.Next
	}
	return result.Next
}