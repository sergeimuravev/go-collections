// Package linkedlist provides single linked list implementation.
package linkedlist

import (
	"../shared"
	"fmt"
)

// Implemented interfaces
var _ shared.Counter = (*LinkedList)(nil)
var _ shared.Enumerable = (*LinkedList)(nil)
var _ fmt.Stringer = (*LinkedList)(nil)

// LinkedList implements single linked list data structure.
type LinkedList struct {
	counter int
	first   *Node
}

// New creates new linked list based on collection of values provided.
func New(values ...interface{}) LinkedList {
	list := LinkedList{}
	if values != nil {
		for _, value := range values {
			list.Add(value)
		}
	}

	return list
}

// Count returns the number of elements in collection.
func (list *LinkedList) Count() int {
	return list.counter
}

// GetEnumerator returns forward iterator implementation.
func (list *LinkedList) GetEnumerator() shared.Enumerator {
	return &enumerator{list, nil}
}

// First returns the head of the list.
func (list *LinkedList) First() *Node {
	return list.first
}

// Add inserts a new node at the beginning of the list.
func (list *LinkedList) Add(value interface{}) *Node {
	newNode := NewNode(value, list.first)
	list.first = &newNode
	list.counter++
	return list.first
}

// Remove deletes node from list.
func (list *LinkedList) Remove(node *Node) bool {
	temp := NewNode(nil, list.first)
	current := &temp
	for node != nil && current != nil {
		if current.next == node {
			current.next = node.next
			list.first = temp.next
			list.counter--
			return true
		}

		current = current.next
	}

	return false
}

// Find returns first node with provided value.
func (list *LinkedList) Find(value interface{}) *Node {
	current := list.first
	for current != nil {
		if current.value == value {
			return current
		}

		current = current.next
	}

	return nil
}

// Clear removes all nodes from the list
func (list *LinkedList) Clear() {
	list.first = nil
	list.counter = 0
}

// String is a Stringer interface implementation.
func (list *LinkedList) String() string {
	return fmt.Sprintf("Linked list. Count=%d.", list.counter)
}
