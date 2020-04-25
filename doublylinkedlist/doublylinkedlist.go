// Package doublylinkedlist provides doubly linked list implementation.
package doublylinkedlist

import (
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*LinkedList)(nil)
var _ shared.Enumerable = (*LinkedList)(nil)

// LinkedList implements doubly linked list data structure.
type LinkedList struct {
	first   *Node
	last    *Node
	counter int
}

// New creates new doubly linked list based on collection of values provided.
// Example: (1,2,3) => 3->2->1
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

	if list.first != nil {
		first := list.first
		first.previous = &node
		list.first = first.previous
	} else {
		// Empty list
		list.first = &node
		list.last = list.first
	}

	list.counter++
	return list.first
}

// AddLast inserts a new node at the end of the list.
func (list *LinkedList) AddLast(value interface{}) *Node {
	node := NewNode(value, nil, list.last)

	if list.last != nil {
		last := list.last
		last.next = &node
		list.last = last.next
	} else {
		// Empty list
		list.last = &node
		list.first = list.last
	}

	list.counter++
	return list.last
}

// AddAfter inserts a new node after the provided one.
// If node is nil, inserts new node at the begining of the list.
func (list *LinkedList) AddAfter(node *Node, value interface{}) *Node {
	if node == nil {
		return list.AddFirst(value)
	}

	newNode := NewNode(value, node.next, node)

	if list.last == newNode.previous {
		list.last = &newNode
	}

	list.counter++
	return &newNode
}

// AddBefore inserts a new node before the provided one.
// If node is nil, inserts new node at the end of the list.
func (list *LinkedList) AddBefore(node *Node, value interface{}) *Node {
	if node == nil {
		return list.AddLast(value)
	}

	newNode := NewNode(value, node, node.previous)

	if list.first == newNode.next {
		list.first = &newNode
	}

	list.counter++
	return &newNode
}

// Remove deletes node from list.
func (list *LinkedList) Remove(node *Node) bool {
	if node == nil {
		return false
	}

	if node.next != nil {
		next := node.next
		next.previous = node.previous
	} else {
		list.last = node.previous // Last node deleted
	}

	if node.previous != nil {
		previous := node.previous
		previous.next = node.next
	} else {
		list.first = node.next // First node deleted
	}

	if node.next != nil || node.previous != nil {
		list.counter--
		return true
	}

	return false
}

// Find returns first node with provided value between from and to nodes. Returns nil if value not found.
// If from node is not provided will start search from the start of the list.
// If to node is not provided will finish search by the end of the list.
func (list *LinkedList) Find(from *Node, to *Node, value interface{}) *Node {
	if list.first == nil {
		return nil // Empty list
	}

	left := from
	if left == nil {
		left = list.first
	}

	right := to
	if right == nil {
		right = list.last
	}

	for left.previous != right &&
		right.next != left {

		if left.value == value {
			return left
		}

		if right.value == value {
			return right
		}

		left = left.next
		right = right.previous

		if left == right {
			if left != nil && left.value == value {
				return left
			}

			break
		}
	}

	return nil
}

// Clear removes all nodes from the list
func (list *LinkedList) Clear() {
	list.first = nil
	list.last = nil
	list.counter = 0
}
