// Package list provides dynamic list implementation which elements can be accessed by index.
package list

import (
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*List)(nil)
var _ shared.Enumerable = (*List)(nil)
var _ shared.Indexer = (*List)(nil)

// List is a dynamic list implementation which elements can be accessed by index.
type List struct {
	buffer []interface{}
}

// New creates list instance from slice. Size and capacity are re-used.
func New(data []interface{}) List {
	return List{buffer: data}
}

// Count returns the number of elements in collection.
func (list *List) Count() int {
	return len(list.buffer)
}

// Capacity returns the internal buffer size.
func (list *List) Capacity() int {
	return cap(list.buffer)
}

// GetEnumerator returns forward iterator implementation.
func (list *List) GetEnumerator() shared.Enumerator {
	return &enumerator{list, -1}
}

// ElementAt returns element by its index.
func (list *List) ElementAt(index int) interface{} {
	return list.buffer[index]
}

// Add inserts new element at the end of the list.
func (list *List) Add(value interface{}) {
	list.buffer = append(list.buffer, value)
}

// AddAll inserts all slice values at the end of the list.
func (list *List) AddAll(values []interface{}) {
	list.buffer = append(list.buffer, values)
}

// Insert inserts new element at the specified index of the list.
func (list *List) Insert(index int, value interface{}) {
	values := make([]interface{}, 1)
	values[0] = value
	list.InsertAll(index, values)
}

// InsertAll inserts all slice values at the specified index of the list.
func (list *List) InsertAll(index int, values []interface{}) {
	tail := make([]interface{}, len(list.buffer)-index)
	copy(tail, list.buffer[:index+1])

	for i := 0; i < len(values); i++ {
		list.buffer[i+index] = values[i]
	}

	list.buffer = append(list.buffer[:index+1+len(values)], tail)
}

// IndexOf returns non-negative index of the first element with provided value or -1 if value not found.
func (list *List) IndexOf(value interface{}) int {
	for i, v := range list.buffer {
		if v == value {
			return i
		}
	}

	return -1
}

// LastIndexOf returns non-negative index of the last element with provided value or -1 if value not found.
func (list *List) LastIndexOf(value interface{}) int {
	for i := len(list.buffer) - 1; i >= 0; i-- {
		if list.buffer[i] == value {
			return i
		}
	}

	return -1
}

// FindAll returns indices of all the elements which values satisfy the provided predicate.
func (list *List) FindAll(predicate func(interface{}) bool) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list.buffer {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

// RemoveAt removes element by its index.
func (list *List) RemoveAt(index int) {
	length := len(list.buffer)
	copy(list.buffer[index:], list.buffer[index+1:])
	list.buffer[length-1] = nil
	list.buffer = list.buffer[:length-1]

	if cap(list.buffer) > 2*length {
		list.TrimExcess() // Shrink buffer
	}
}

// RemoveAll removes all elements in the list with the value provided.
func (list *List) RemoveAll(value interface{}) {
	for i := list.IndexOf(value); i >= 0; {
		list.RemoveAt(i)
	}
}

// TrimExcess sets the capacity to the actual number of elements in the list plus one.
func (list *List) TrimExcess() {
	buffer := make([]interface{}, len(list.buffer)+1)
	copy(buffer, list.buffer)
	list.buffer = buffer
}
