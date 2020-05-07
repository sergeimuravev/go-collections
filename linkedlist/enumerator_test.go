package linkedlist_test

import (
	. "github.com/sergeimuravev/go-collections/linkedlist"
	"testing"
)

func TestMoveNext(t *testing.T) {
	samples := []struct {
		list   LinkedList
		values []int
	}{
		{LinkedList{}, make([]int, 0)},
		{New(1, 2, 3), []int{3, 2, 1}},
		{New(0, 0, 0), []int{0, 0, 0}},
		{New(3, 2, 1), []int{1, 2, 3}},
	}

	for _, sample := range samples {
		it := sample.list.GetEnumerator()

		var nodes []*Node
		for it.MoveNext() {
			nodes = append(nodes, it.Current().(*Node))
		}

		for i, node := range nodes {
			if node.Value() != sample.values[i] {
				t.Errorf("Incorrect value, got: %d, expected: %d.", node.Value(), sample.values[i])
			}
		}
	}
}

func TestReset(t *testing.T) {
	samples := []struct {
		list  LinkedList
		value interface{}
	}{
		{LinkedList{}, nil},
		{New(1, 2, 3), 3},
		{New(0, 0, 0), 0},
		{New(3, 2, 1), 1},
	}

	for _, sample := range samples {
		it := sample.list.GetEnumerator()
		it.MoveNext()
		it.Reset()

		if it.MoveNext() {
			node := it.Current().(*Node)
			if node.Value() != sample.value {
				t.Errorf("Incorrect value, got: %d, expected: %d.", node.Value(), sample.value)
			}
		}
	}
}
