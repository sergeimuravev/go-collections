// Package avltree provide Adelson-Velsky/Landis balanced binary search tree implementation.
// Balanced tree guarantees logarithmic time complexity for major operations.
package avltree

import (
	"github.com/sergeimuravev/go-collections/shared"
)

// Implemented interfaces
var _ shared.Counter = (*Tree)(nil)
var _ shared.Enumerable = (*Tree)(nil)
var _ shared.Dictionary = (*Tree)(nil)

// Tree is an AVL tree implementation.
type Tree struct {
	root     *Node
	count    int
	comparer shared.Comparer
}

// New constructs empty tree.
func New(comparer shared.Comparer) Tree {
	return Tree{comparer: comparer}
}

// Count returns the number of elements in collection.
func (tree *Tree) Count() int {
	return tree.count
}

// Add inserts a new node to tree.
func (tree *Tree) Add(key, value interface{}) {
	tree.root = tree.insert(tree.root, key, value)
	tree.count++
}

// Remove deleted node from tree.
// Returns true if element with provided key exists
// in collection, otherwise returns false.
func (tree *Tree) Remove(key interface{}) bool {
	node, removed := tree.remove(tree.root, key)
	tree.root = node

	if removed {
		tree.count--
	}

	return removed
}

// ContainsKey returns true if collection contains an element with provided key.
func (tree *Tree) ContainsKey(key interface{}) bool {
	return tree.search(tree.root, key) != nil
}

// Get returns value by key.
func (tree *Tree) Get(key interface{}) interface{} {
	node := tree.search(tree.root, key)
	if node == nil {
		return nil
	}

	return node.Value
}

// GetEnumerator returns forward iterator implementation.
func (tree *Tree) GetEnumerator() shared.Enumerator {
	return &enumerator{tree: tree}
}

// Clear removes all nodes from the tree.
func (tree *Tree) Clear() {
	tree.root = nil
	tree.count = 0
}

// compare checks if tree comparer assigned and uses default one
// if no one was provided
func (tree *Tree) compare(left, right interface{}) int {
	cmp := tree.comparer
	if cmp == nil {
		// Default comparer
		cmp = func(left, right interface{}) int {
			if left == right {
				return 0
			}

			return 1
		}
	}

	return cmp(left, right)
}

// insert recursively adds new node into tree
func (tree *Tree) insert(p *Node, key, value interface{}) *Node {
	if p == nil {
		newNode := Node{Key: key, Value: value, height: 1}
		return &newNode
	}

	if tree.compare(key, p.Key) < 0 {
		p.left = tree.insert(p.left, key, value)
	} else {
		p.right = tree.insert(p.right, key, value)
	}

	return p.balance()
}

// remove recursively removes node with provided key from tree
func (tree *Tree) remove(p *Node, key interface{}) (*Node, bool) {
	if p == nil {
		return nil, false
	}

	removed := false
	if tree.compare(key, p.Key) < 0 {
		p.left, removed = tree.remove(p.left, key)
	} else if tree.compare(key, p.Key) > 0 {
		p.right, removed = tree.remove(p.right, key)
	} else {
		q := p.left
		r := p.right
		if r == nil {
			return q, true
		}

		min := r.findMin()
		min.right = r.removeMin()
		min.left = q
		return min.balance(), true
	}

	return p.balance(), removed
}

// search performs DFS in the binary search tree
func (tree *Tree) search(p *Node, key interface{}) *Node {
	if p == nil {
		return nil
	}

	if tree.compare(key, p.Key) < 0 {
		return tree.search(p.left, key)
	} else if tree.compare(key, p.Key) > 0 {
		return tree.search(p.right, key)
	}
	return p
}
