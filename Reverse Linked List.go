//206. Reverse Linked List
//https://leetcode.com/problems/reverse-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	var nextTemp *ListNode
	for curr != nil {
		nextTemp = curr.Next //2>3>4
		curr.Next = prev     // nil>1>2
		prev = curr          // 1>2>3
		curr = nextTemp      //2>3>4
	}
	return prev
}