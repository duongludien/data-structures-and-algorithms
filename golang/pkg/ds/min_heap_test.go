package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeapify(t *testing.T) {
	arr := []int{5, 6, 2, 2, 10, 12, 9, 10, 9, 3}
	heap, error := Heapify(arr)

	assert.Nil(t, error)
	assert.NotNil(t, heap)
	assert.NotEqual(t, 0, len(heap))

	for i := 0; i <= (len(heap)-1)/2; i++ {
		assert.True(t, heap[i] <= heap[2*i+1])
		if 2*i+2 < len(heap) {
			assert.True(t, heap[i] <= heap[2*i+2])
		}
	}
}
