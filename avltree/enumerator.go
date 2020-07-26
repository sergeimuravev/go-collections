package avltree

import (
	"github.com/sergeimuravev/go-collections/queue"
	"github.com/sergeimuravev/go-collections/shared"
)

// Implemented interfaces
var _ shared.Enumerator = (*enumerator)(nil)

const bufferSize = 8

type enumerator struct {
	tree    *Tree
	current *Node
	buffer  queue.Queue
}

func (it *enumerator) Current() interface{} {
	if it.current == nil {
		return nil
	}

	return it.current
}

func (it *enumerator) MoveNext() bool {
	if it.current == nil {
		it.current = it.tree.root
		return it.current != nil
	}

	if it.buffer.Count() == 0 {
		it.loadBuffer(it.current, false)
	}

	if it.buffer.Count() > 0 {
		p := it.buffer.Dequeue()
		it.current = p.(*Node)
		return true
	}

	return false
}

func (it *enumerator) Reset() {
	it.current = nil
}

// Pre-order tree traversal to populate buffer
func (it *enumerator) loadBuffer(node *Node, enqueue bool) {
	if node == nil ||
		it.buffer.Count() >= bufferSize {
		return
	}

	if enqueue {
		it.buffer.Enqueue(node)
	}

	if node.left != nil {
		it.loadBuffer(node.left, true)
	}

	if node.right != nil {
		it.loadBuffer(node.right, true)
	}
}
