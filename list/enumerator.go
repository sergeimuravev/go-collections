package list

import (
	"../shared"
)

// Implemented interfaces
var _ shared.Enumerator = (*enumerator)(nil)

type enumerator struct {
	list  *List
	index int
}

func (it *enumerator) Current() interface{} {
	if it.index < 0 {
		return nil
	}

	return it.list.ElementAt(it.index)
}

func (it *enumerator) MoveNext() bool {
	if it.index < 0 {
		if it.list.Count() == 0 {
			return false // Empty list
		}
	}

	it.index++
	return it.index >= 0
}

func (it *enumerator) Reset() {
	it.index = -1
}
