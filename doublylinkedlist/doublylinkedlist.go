package doublylinkedlist

import (
	"../shared"
)

// LinkedList implements doubly linked list data structure.
type LinkedList struct {
	first   *Node
	last    *Node
	counter int
}

// New creates new doubly linked list based on collection of values provided.
func New(values ...interface{}) LinkedList {
	list := LinkedList{}
	var next *Node
	if values != nil {
		for _, value := range values {
			node := Node{value: value, next: next}

			if next == nil {
				list.last = &node
			} else {
				next.previous = &node
			}

			next = &node
		}

		list.first = next
		next.previous = nil
		list.counter++
	}

	return list
}

// GetEnumerator returns forward iterator implementation.
func (list *LinkedList) GetEnumerator() shared.Enumerator {
	return &enumerator{list, nil}
}

// Count returns the number of elements in collection.
func (list *LinkedList) Count() int {
	return list.counter
}

// First returns the head of the list.
func (list *LinkedList) First() *Node {
	return list.first
}

// Last returns the tail of the list.
func (list *LinkedList) Last() *Node {
	return list.last
}
