package queues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateANew(t *testing.T) {
	queue := NewQueue()

	assert.NotNil(t, queue, "Should not be nil")
	assert.Equal(t, []int{0}, queue.Data)
	assert.Equal(t, 0, queue.Length)
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(t, 3, queue.Length, "Should add one new element to the queue")
	expected := []int{1, 2, 3}
	assert.Equal(t, expected, queue.Data)
}

func TestDequeue(t *testing.T) {
	queue := NewQueue()

	expected := []int([]int(nil))
	queue.Dequeue()
	assert.Equal(t, expected, queue.Data)
	assert.Equal(t, 0, queue.Length, "Should remove one element from the queue")

	queue.Enqueue(1).Enqueue(2).Enqueue(3).Enqueue(4)
	queue.Dequeue()

	expected = []int{2, 3, 4}
	assert.Equal(t, expected, queue.Data)
	assert.Equal(t, 3, queue.Length, "Should remove one element from the queue")
}

func TestPeek(t *testing.T) {
	queue := NewQueue()

	expected := 0
	assert.Equal(t, expected, queue.Peek(), "Should return 0 when the queue is empty")

	queue.Enqueue(1)
	expected = 1
	assert.Equal(t, expected, queue.Peek(), "Should return the first element in the queue")
}
