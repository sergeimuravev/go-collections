package list

import (
	"github.com/sergeimuravev/go-collections/shared"
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
	count := it.list.Count()

	if count == 0 {
		return false // Empty list
	}

	index := it.index
	if it.index < count-1 {
		it.index++
	}

	return index != it.index
}

func (it *enumerator) Reset() {
	it.index = -1
}
