// Package list provides dynamic array implementation.
package list

import (
	"../shared"
)

// Implemented interfaces
var _ shared.Counter = (*List)(nil)
var _ shared.Enumerable = (*List)(nil)
var _ shared.Indexer = (*List)(nil)

type List struct {
	buffer []interface{}
}

func New(data []interface{}) List {
	return List{buffer: data}
}

// Count returns the number of elements in collection.
func (list *List) Count() int {
	return len(list.buffer)
}

func (list *List) Capacity() int {
	return cap(list.buffer)
}

// GetEnumerator returns forward iterator implementation.
func (list *List) GetEnumerator() shared.Enumerator {
	return &enumerator{list, -1}
}

func (list *List) ElementAt(index int) interface{} {
	return list.buffer[index]
}

func (list *List) Add(value interface{}) {
	list.buffer = append(list.buffer, value)
}

func (list *List) AddAll(values []interface{}) {
	list.buffer = append(list.buffer, values)
}

func (list *List) IndexOf(value interface{}) int {
	for i, v := range list.buffer {
		if v == value {
			return i
		}
	}

	return -1
}

func (list *List) LastIndexOf(value interface{}) int {
	for i := len(list.buffer) - 1; i >= 0; i-- {
		if list.buffer[i] == value {
			return i
		}
	}

	return -1
}

func (list *List) FindAll(predicate func(interface{}) bool) []interface{} {
	result := make([]interface{}, 0)
	for _, v := range list.buffer {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func (list *List) RemoveAt(index int) {
	length := len(list.buffer)
	copy(list.buffer[index:], list.buffer[index+1:])
	list.buffer[length-1] = nil
	list.buffer = list.buffer[:length-1]

	if cap(list.buffer) > 2*length {
		list.TrimExcess() // Shrink buffer
	}
}

func (list *List) RemoveAll(value interface{}) {
	for i := list.IndexOf(value); i >= 0; {
		list.RemoveAt(i)
	}
}

func (list *List) TrimExcess() {
	buffer := make([]interface{}, len(list.buffer)+1)
	copy(buffer, list.buffer)
	list.buffer = buffer
}
