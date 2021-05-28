//143. Reorder List
//https://leetcode.com/problems/reorder-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodeList, root := make([]*ListNode, 0), head
	for root != nil {
		nodeList = append(nodeList, root)
		root = root.Next
	}
	l, i := len(nodeList), 0
	for (l%2 == 0 && i < (l-1)/2) || (l%2 == 1 && i < l/2) {
		nodeList[i].Next = nodeList[l-i-1]
		nodeList[l-i-1].Next = nodeList[i+1]
		nodeList[l-i-2].Next = nil
		i++
	}
}