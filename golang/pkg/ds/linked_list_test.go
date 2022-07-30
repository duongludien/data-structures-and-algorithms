package ds

import "testing"

func TestInitFirstNode(t *testing.T) {
	list := NewListNode(3)

	if list.Value != 3 {
		t.Fatalf(`list.Value = %d, want %d`, list.Value, 3)
	}

	if list.Next != nil {
		t.Fatalf(`list.Next = %p, want nil`, list.Next)
	}
}

func TestInsertNode(t *testing.T) {
	list := NewListNode(0)

	node := NewListNode(4)
	list.InsertNode(node)

	if list.Count() != 1 {
		t.Fatalf(`list.Count() = %d, want 1`, list.Count())
	}

	if list.Next.Value != 4 {
		t.Fatalf(`lsit.Next.Value = %d, want 4`, list.Next.Value)
	}
}

func TestRemoveNode(t *testing.T) {
	list := NewListNode(0)

	node1 := NewListNode(4)
	list.InsertNode(node1)

	node2 := NewListNode(5)
	node1.InsertNode(node2)

	node3 := NewListNode(6)
	node2.InsertNode(node3)

	if list.Count() != 3 {
		t.Fatalf(`Initial state: list.Count() = %d, want 3`, list.Count())
	}

	list.RemoveNode(node2)

	if list.Count() != 2 {
		t.Fatalf(`After removing: list.Count() = %d, want 2`, list.Count())
	}

	if list.Next.Value != 4 {
		t.Fatalf(`First node value = %d, want 3`, list.Next.Value)
	}

	if list.Next.Next.Value != 6 {
		t.Fatalf(`Second node value = %d, want 5`, list.Next.Next.Value)
	}
}
