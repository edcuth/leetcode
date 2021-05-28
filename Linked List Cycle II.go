//142. Linked List Cycle II
//https://leetcode.com/problems/linked-list-cycle-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodes := make(map[*ListNode]bool)
	cur := head
	for cur.Next != nil {
		if _, ok := nodes[cur.Next]; ok {
			return cur.Next
		} else {
			nodes[cur] = true
			cur = cur.Next
		}
	}
	return cur.Next
}