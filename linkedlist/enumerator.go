package linkedlist

import (
	"github.com/sergeimuravev/go-collections/shared"
)

// Implemented interfaces
var _ shared.Enumerator = (*enumerator)(nil)

type enumerator struct {
	list    *LinkedList
	current *Node
}

func (it *enumerator) Current() interface{} {
	if it.current == nil {
		return nil
	}

	return it.current
}

func (it *enumerator) MoveNext() bool {
	if it.current == nil {
		if it.list.first == nil {
			return false // Empty list
		}

		it.current = it.list.first
	} else if it.current.next != nil {
		it.current = it.current.next
	} else {
		return false
	}

	return it.current != nil
}

func (it *enumerator) Reset() {
	it.current = nil
}
