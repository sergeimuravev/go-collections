// Package heap provides binary heap implementation.
package heap

import (
	"github.com/sergeimuravev/go-collections/list"
	"github.com/sergeimuravev/go-collections/shared"
)

// Implemented interfaces
var _ shared.Counter = (*Heap)(nil)
var _ shared.Enumerable = (*Heap)(nil)

// Ascending order comparator: returns true if the left value is less than the right one.
type comparer func(interface{}, interface{}) bool

// Heap represents a binary heap.
type Heap struct {
	list list.List // [root, 1, 2, 11, 12, 21, 22]
	cmp  comparer
}

// New creates heap instance from slice.
func New(data []interface{}, cmp comparer) Heap {
	list := list.New(make([]interface{}, 0, cap(data)))
	if cmp == nil {
		cmp = func(x interface{}, y interface{}) bool {
			return x == y
		}
	}

	heap := Heap{list, cmp}
	for _, v := range data {
		heap.Push(v)
	}

	return heap
}

// Count returns the number of elements in collection.
func (heap *Heap) Count() int {
	return heap.list.Count()
}

// GetEnumerator returns forward iterator implementation.
func (heap *Heap) GetEnumerator() shared.Enumerator {
	return heap.list.GetEnumerator()
}

// Push adds an element to the collection.
func (heap *Heap) Push(value interface{}) {
	list := &heap.list
	list.Add(value)
	// Swim
	k := list.Count() - 1
	for k > 0 && heap.cmp(list.ElementAt(k/2), list.ElementAt(k)) {
		heap.exchange(k/2, k)
		k = k / 2
	}
}

// Pop returns top element of the heap.
func (heap *Heap) Pop() interface{} {
	list := &heap.list
	if list.Count() == 0 {
		return nil
	}
	// Swap top and bottom, remove top
	top := list.ElementAt(0)
	index := list.Count() - 1
	heap.exchange(0, index)
	list.RemoveAt(index)
	// Sink
	k := 0
	for 2*k < list.Count() {
		j := 2 * k

		if j < list.Count()-1 && heap.cmp(list.ElementAt(j), list.ElementAt(j+1)) {
			j++
		}

		if !heap.cmp(list.ElementAt(k), list.ElementAt(j)) {
			break
		}

		heap.exchange(k, j)
		k = j
	}

	return top
}

// Peek returns top element of the heap.
func (heap *Heap) Peek() interface{} {
	if heap.Count() == 0 {
		return nil
	}

	return heap.list.ElementAt(0)
}

// Exchange swaps values at specified positions
func (heap *Heap) exchange(left, right int) {
	list := heap.list
	lValue := list.ElementAt(left)
	rValue := list.ElementAt(right)
	list.SetElementAt(left, rValue)
	list.SetElementAt(right, lValue)
}
