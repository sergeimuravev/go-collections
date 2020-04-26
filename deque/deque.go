// Package deque implements double-ended queue abstract data type.
package deque

import (
	"../doublylinkedlist"
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*Deque)(nil)
var _ shared.Enumerable = (*Deque)(nil)

// Deque is a double-ended queue data structure implementation.
type Deque struct {
	list *doublylinkedlist.LinkedList
}

// New creates new deque based on collection of values provided.
func New(values ...interface{}) Deque {
	list := doublylinkedlist.New(values)
	deque := Deque{&list}
	return deque
}

// Count returns the number of elements in collection.
func (deque *Deque) Count() int {
	return deque.list.Count()
}

// GetEnumerator returns forward iterator implementation.
func (deque *Deque) GetEnumerator() shared.Enumerator {
	return deque.list.GetEnumerator()
}

// Front returns the first element in the deque.
func (deque *Deque) Front() interface{} {
	node := deque.list.First()
	if node == nil {
		return nil
	}

	return node.Value()
}

// Back returns the last element in the deque.
func (deque *Deque) Back() interface{} {
	node := deque.list.Last()
	if node == nil {
		return nil
	}

	return node.Value()
}

// PushFront adds an element to the beginning of the deque.
func (deque *Deque) PushFront(value interface{}) {
	deque.list.AddFirst(value)
}

// PushBack adds an element to the end of the deque.
func (deque *Deque) PushBack(value interface{}) {
	deque.list.AddLast(value)
}

// PopFront removes and returns the first element of the deque.
func (deque *Deque) PopFront() interface{} {
	node := deque.list.First()
	if deque.list.Remove(node) {
		return node.Value()
	}

	return nil
}

// PopBack removes and returns the last element of the deque.
func (deque *Deque) PopBack() interface{} {
	node := deque.list.Last()
	if deque.list.Remove(node) {
		return node.Value()
	}

	return nil
}

// Clear removes all entries from deque.
func (deque *Deque) Clear() {
	deque.list.Clear()
}
