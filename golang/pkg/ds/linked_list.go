package ds

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(value int) *ListNode {
	node := new(ListNode)
	node.Value = value
	node.Next = nil

	return node
}

func (listNode *ListNode) InsertNode(anotherNode *ListNode) {
	anotherNode.Next = listNode.Next
	listNode.Next = anotherNode
}

func (list *ListNode) Count() int {
	iterator := list
	count := 0

	for iterator.Next != nil {
		count++
		iterator = iterator.Next
	}

	return count
}

func (list *ListNode) RemoveNode(targetNode *ListNode) {
	iterator := list

	for iterator.Next != nil {
		if iterator.Next == targetNode {
			iterator.Next = iterator.Next.Next
			iterator.Next.Next = nil
			break
		}

		iterator = iterator.Next
	}
}
