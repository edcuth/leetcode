//138. Copy List with Random Pointer
//https://leetcode.com/problems/copy-list-with-random-pointer/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	/* var result is the head for a new linked list.
	 * p is a pointer we'll be using to move through result.
	 * temp is a pointer to the head of the linked list we need to copy.
	 * map nodes will be used later to add pointers to the Random field.
	 */
	var (
		result *Node           = &Node{}
		p                      = result
		temp   *Node           = head
		nodes  map[*Node]*Node = make(map[*Node]*Node, 0)
	)

	/* this loop creates a new Node and adds a pointer to the Next field of the node in our p pointer.
	 * then we add the temp pointer as key for our p.Next to our nodes map.
	 * we advance the p pointer to the next node in the list (the one we created at the beginning of the loop).
	 * we advance our temp pointer to the next in the list.
	 */
	for temp != nil {
		p.Next = &Node{Val: temp.Val}
		nodes[temp] = p.Next
		p = p.Next
		temp = temp.Next
	}

	/* we reset our temp and p pointers to point to the beginning of both linked lists again */
	temp = head
	p = result.Next

	/* And this final for loop will use your nodes map to add pointers to the Random field of our new copied list
	 * then advance the pointers! */
	for temp != nil {
		if n, found := nodes[temp.Random]; found {
			p.Random = n
		}
		p = p.Next
		temp = temp.Next
	}
	return result.Next
}