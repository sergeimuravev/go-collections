// Package linkedlist provides singly linked list implementation.
package linkedlist

import (
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*LinkedList)(nil)
var _ shared.Enumerable = (*LinkedList)(nil)

// LinkedList implements singly linked list data structure.
type LinkedList struct {
	counter int
	first   *Node
}

// New creates new linked list based on collection of values provided.
func New(values ...interface{}) LinkedList {
	list := LinkedList{}
	if values != nil {
		for _, value := range values {
			if slice, ok := value.([]interface{}); len(values) == 1 && ok {
				for _, elem := range slice {
					list.Add(nil, elem) // If slice of values provided
				}
			} else {
				list.Add(nil, value)
			}
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

// Add inserts a new node after the provided one.
// If node is not provided, inserts a new node at the beginning of the list.
func (list *LinkedList) Add(node *Node, value interface{}) *Node {
	var newNode *Node
	if node == nil {
		newNode = &Node{value: value, next: list.first}
		list.first = newNode
	} else {
		newNode = &Node{value: value, next: node.next}
		node.next = newNode
	}

	list.counter++
	return newNode
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

// Find returns first node with provided value staring from the given node. Returns nil if value not found.
// If node is not provided will start search from the beginning of the list.
func (list *LinkedList) Find(node *Node, value interface{}) *Node {
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

// Clear removes all nodes from the list.
func (list *LinkedList) Clear() {
	list.first = nil
	list.counter = 0
}
