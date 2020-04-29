// Package queue implements queue abstract data type.
package queue

import (
	"../linkedlist"
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*Queue)(nil)
var _ shared.Enumerable = (*Queue)(nil)

// Queue is FIFO data structure implementation.
type Queue struct {
	list linkedlist.LinkedList
	tail *linkedlist.Node
}

// New creates new queue based on collection of values provided.
func New(values ...interface{}) Queue {
	queue := Queue{linkedlist.LinkedList{}, nil}
	if values != nil {
		var prev *linkedlist.Node
		for _, value := range values {
			if slice, ok := value.([]interface{}); len(values) == 1 && ok {
				for _, elem := range slice {
					prev = queue.list.Add(prev, elem) // If slice of values provided
				}
			} else {
				prev = queue.list.Add(prev, value)
			}
		}
	}

	return queue
}

// Count returns the number of elements in collection.
func (queue *Queue) Count() int {
	return queue.list.Count()
}

// GetEnumerator returns forward iterator implementation.
func (queue *Queue) GetEnumerator() shared.Enumerator {
	return queue.list.GetEnumerator()
}

// Enqueue adds value to the end of the queue.
func (queue *Queue) Enqueue(value interface{}) {
	node := queue.list.Add(queue.tail, value)
	queue.tail = node
}

// Dequeue removes and returns value from the begining of the queue.
func (queue *Queue) Dequeue() interface{} {
	first := queue.list.First()
	if queue.list.Remove(first) {
		if first == queue.tail {
			queue.tail = nil // Empty queue
		}

		return first.Value()
	}

	return nil
}

// Peek returns value from the begining of the queue.
func (queue *Queue) Peek() interface{} {
	node := queue.list.First()
	if node == nil {
		return nil
	}

	return node.Value()
}

// Clear removes all entries from queue.
func (queue *Queue) Clear() {
	queue.tail = nil
	queue.list.Clear()
}
