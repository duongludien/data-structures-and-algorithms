package ds

type ListNode struct {
    Value int
    Next  *ListNode
}

func NewListNode(value int) *ListNode {
    return &ListNode{Value: value}
}

func BuildList(arr []int) *ListNode {
    if len(arr) == 0 {
        return nil
    }

    head := NewListNode(arr[0])
    iterator := head
    for i := 1; i < len(arr); i++ {
        iterator.Next = NewListNode(arr[i])
        iterator = iterator.Next
    }

    return head
}

func (head *ListNode) GetLength() int {
    iterator := head
    length := 0

    for iterator != nil {
        length++
        iterator = iterator.Next
    }

    return length
}

func (head *ListNode) Traverse() {
    iterator := head

    for iterator != nil {
        PrintInt(iterator.Value)
        iterator = iterator.Next
    }
}

func (head *ListNode) Push(value int) *ListNode {
    node := NewListNode(value)

    if head.GetLength() == 0 {
        return node
    }

    iterator := head

    for iterator.Next != nil {
        iterator = iterator.Next
    }

    iterator.Next = node

    return head
}

func (head *ListNode) Pop() (*ListNode, int) {
    list := head
    lastValue := 0

    if list.GetLength() == 0 {
        // Do nothing
    } else if list.GetLength() == 1 {
        lastValue = list.Value
        list = nil
    } else {
        iterator := list.Next
        previousNode := list
        for iterator.Next != nil {
            previousNode = iterator
            iterator = iterator.Next
        }

        lastValue = previousNode.Next.Value
        previousNode.Next = nil
    }

    return list, lastValue
}
