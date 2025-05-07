package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type CircularQueue struct {
	values      []int
	startIdx    int
	sizeOfQueue int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{values: make([]int, size), startIdx: 0, sizeOfQueue: 0}
}

func (q *CircularQueue) Push(value int) bool {
	if q.sizeOfQueue == len(q.values) {
		return false
	}
	q.sizeOfQueue++
	q.values[(q.startIdx+q.sizeOfQueue-1)%len(q.values)] = value
	return true
}

func (q *CircularQueue) Pop() bool {
	if q.Empty() {
		return false
	}
	q.startIdx = (q.startIdx + 1) % len(q.values)
	q.sizeOfQueue--
	return true
}

func (q *CircularQueue) Front() int {
	if q.Empty() {
		return -1
	}
	return q.values[q.startIdx]
}

func (q *CircularQueue) Back() int {
	if q.Empty() {
		return -1
	}
	return q.values[(q.startIdx+q.sizeOfQueue-1)%len(q.values)]
}

func (q *CircularQueue) Empty() bool {
	return q.sizeOfQueue == 0
}

func (q *CircularQueue) Full() bool {
	return q.sizeOfQueue == len(q.values)
}

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.values))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
