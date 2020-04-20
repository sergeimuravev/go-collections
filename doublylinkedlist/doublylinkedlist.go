// Package doublylinkedlist provides doubly linked list implementation.
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
	if values != nil {
		for _, value := range values {
			list.AddFirst(value)
		}
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

// AddFirst inserts a new node at the begining of the list.
func (list *LinkedList) AddFirst(value interface{}) *Node {
	node := NewNode(value, list.first, nil)
	list.first = &node
	list.counter++
	return list.first
}

// AddLast inserts a new node at the end of the list.
func (list *LinkedList) AddLast(value interface{}) *Node {
	node := NewNode(value, nil, list.last)
	list.last = &node
	list.counter++
	return list.last
}

// AddAfter inserts a new node after the provided one.
func (list *LinkedList) AddAfter(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	newNode := NewNode(value, node.next, node)
	list.counter++
	return &newNode
}

// AddBefore inserts a new node before the provided one.
func (list *LinkedList) AddBefore(node *Node, value interface{}) *Node {
	if node == nil {
		return nil
	}

	newNode := NewNode(value, node, node.previous)
	list.counter++
	return &newNode
}

// Remove deletes node from list.
func (list *LinkedList) Remove(node *Node) bool {
	if node.next != nil {
		next := *node.next
		next.previous = node.previous
	}

	if node.previous != nil {
		previous := *node.previous
		previous.next = node.next
	}

	if node.next != nil || node.previous != nil {
		list.counter--
		return true
	}

	return false
}

// FindAfter returns first node with provided value staring after the given node. Returns nil if value not found.
// If node is not provided will start search from the beginning of the list.
func (list *LinkedList) FindAfter(node *Node, value interface{}) *Node {
	current := node
	if node == nil {
		current = list.first
	}

	for current != nil {
		if current.value == value {
			return current
		}

		current = current.next
	}

	return nil
}

// FindBefore returns first node with provided value staring before the given node. Returns nil if value not found.
// If node is not provided will start search from the end of the list.
func (list *LinkedList) FindBefore(node *Node, value interface{}) *Node {
	current := node
	if node == nil {
		current = list.last
	}

	for current != nil {
		if current.value == value {
			return current
		}

		current = current.previous
	}

	return nil
}

// Clear removes all nodes from the list
func (list *LinkedList) Clear() {
	list.first = nil
	list.last = nil
	list.counter = 0
}
