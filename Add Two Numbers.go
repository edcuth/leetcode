// Add Two Numbers
//https://leetcode.com/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{Val: 0}
	i := l1
	j := l2
	current := result
	var carry int = 0
	for i != nil || j != nil {
		var x int = 0
		var y int = 0
		var sum int = 0
		if i != nil {
			x = i.Val
		}
		if j != nil {
			y = j.Val
		}
		sum += x + y + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		if i != nil {
			i = i.Next
		}
		if j != nil {
			j = j.Next
		}
		if carry > 0 {
			current.Next = &ListNode{Val: carry}
		}

	}
	return result.Next
}