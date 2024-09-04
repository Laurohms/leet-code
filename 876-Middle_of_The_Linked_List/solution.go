/*
	Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.
*/
package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	lNode := make([]*ListNode, 0, 10)
	for head != nil {
		lNode = append(lNode, head)
		head = head.Next
	}
	return lNode[len(lNode)/2]
}
