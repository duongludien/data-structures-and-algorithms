package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewListNode_HasCorrectValueAndNextIsNil(t *testing.T) {
	listNodeValue := 2
	head := NewListNode(listNodeValue)
	assert.Equal(t, listNodeValue, head.Value)
	assert.Nil(t, head.Next)
}

func Test_BuildList_PassNil_NilList(t *testing.T) {
	head := BuildList(nil)
	assert.Nil(t, head)
}

func Test_BuildList_PassEmptyArray_NilList(t *testing.T) {
	head := BuildList([]int{})
	assert.Nil(t, head)
}

func Test_BuildList_PassOneItemArray_OneItemList(t *testing.T) {
	head := BuildList([]int{1})
	assert.Equal(t, 1, head.GetLength())

	_, value := head.Pop()
	assert.Equal(t, 1, value)
}

func Test_BuildList_PassLongArray_CorrectList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	length := len(arr)

	head := BuildList(arr)
	assert.Equal(t, length, head.GetLength())

	poppedValue := 0
	for index := length - 1; index >= 0; index-- {
		head, poppedValue = head.Pop()
		assert.Equal(t, arr[index], poppedValue)
	}
}

func Test_Traverse(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5})
	head.Traverse()
}

func Test_Push_EmptyList_OneItemList(t *testing.T) {
	valueToPush := 2
	var head *ListNode = nil
	assert.Zero(t, head.GetLength())

	newList := head.Push(valueToPush)
	assert.Equal(t, 1, newList.GetLength())

	_, lastValue := newList.Pop()
	assert.Equal(t, valueToPush, lastValue)
}

func Test_Push_OneItemList_TwoItemsList(t *testing.T) {
	valueToPush := 2
	var head *ListNode = NewListNode(3)
	assert.Equal(t, 1, head.GetLength())

	newList := head.Push(valueToPush)
	assert.Equal(t, 2, newList.GetLength())

	_, lastValue := newList.Pop()
	assert.Equal(t, valueToPush, lastValue)
}

func Test_Push_LongList_LongerList(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 5, head.GetLength())

	firstValue := 6
	secondValue := 7

	head = head.Push(firstValue).Push(secondValue)
	assert.Equal(t, 5+2, head.GetLength())

	_, lastValue := head.Pop()
	assert.Equal(t, secondValue, lastValue)

	_, lastValue = head.Pop()
	assert.Equal(t, firstValue, lastValue)
}

func Test_Pop_OneItemList_EmptyList(t *testing.T) {
	itemValue := 2
	head := NewListNode(itemValue)

	newList, lastValue := head.Pop()

	assert.Zero(t, newList.GetLength())
	assert.Equal(t, itemValue, lastValue)
}

func Test_Pop_EmptyList_DoNothing(t *testing.T) {
	var head *ListNode = nil

	newList, lastValue := head.Pop()

	assert.Zero(t, newList.GetLength())
	assert.Zero(t, lastValue)
}

func Test_Pop_TwoElements_LastItemRemoved(t *testing.T) {
	head := BuildList([]int{2, 3})
	assert.Equal(t, 2, head.GetLength())

	newList, lastValue := head.Pop()
	assert.Equal(t, 1, newList.GetLength())
	assert.Equal(t, 3, lastValue)
}

func Test_Pop_MoreElements_ChainingRemove(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 6, head.GetLength())

	newList, lastValue := head.Pop()
	assert.Equal(t, 5, newList.GetLength())
	assert.Equal(t, 6, lastValue)

	newList, lastValue = newList.Pop()
	newList, lastValue = newList.Pop()
	assert.Equal(t, 3, newList.GetLength())

	newList, lastValue = newList.Pop()
	newList, lastValue = newList.Pop()
	newList, lastValue = newList.Pop()
	assert.Equal(t, 0, newList.GetLength())
}
