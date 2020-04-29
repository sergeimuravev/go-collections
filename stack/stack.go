// Package stack implements stack abstract data type.
package stack

import (
	"../linkedlist"
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*Stack)(nil)
var _ shared.Enumerable = (*Stack)(nil)

// Stack is LIFO data structure implementation.
type Stack struct {
	list linkedlist.LinkedList
}

// New creates new stack based on collection of values provided.
func New(values ...interface{}) Stack {
	list := linkedlist.New(values)
	stack := Stack{list}
	return stack
}

// Count returns the number of elements in collection.
func (stack *Stack) Count() int {
	return stack.list.Count()
}

// GetEnumerator returns forward iterator implementation.
func (stack *Stack) GetEnumerator() shared.Enumerator {
	return stack.list.GetEnumerator()
}

// Push adds an element to the collection.
func (stack *Stack) Push(value interface{}) {
	stack.list.Add(nil, value)
}

// Pop removes and returns the most recently added element.
func (stack *Stack) Pop() interface{} {
	node := stack.list.First()
	if node == nil {
		return nil
	}

	if !stack.list.Remove(node) {
		return nil
	}

	return node.Value()
}

// Peek returns the most recently added element.
func (stack *Stack) Peek() interface{} {
	node := stack.list.First()
	if node == nil {
		return nil
	}

	return node.Value()
}

// Clear removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.list.Clear()
}
