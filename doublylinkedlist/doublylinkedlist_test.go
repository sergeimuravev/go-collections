package doublylinkedlist_test

import (
	. "../doublylinkedlist"
	"testing"
)

func TestNew(t *testing.T) {
	samples := []struct {
		list   LinkedList
		values []int
	}{
		{New(1, 2, 3, 4, 5), []int{5, 4, 3, 2, 1}},
		{New(5, 4, 3, 2, 1), []int{1, 2, 3, 4, 5}},
		{New(), make([]int, 0)},
	}

	for _, sample := range samples {
		length := len(sample.values)
		if sample.list.Count() != length {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), length)
		}

		var prev, next *Node
		it := sample.list.GetEnumerator()
		for i := 0; i < length; i++ {
			if !it.MoveNext() {
				break
			}

			node := it.Current().(*Node)
			current := node.Value()
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect node value, got: %d, expected: %d.", current, expected)
			}

			if prev != nil {
				if prev.Value() != node.Previous().Value() {
					t.Errorf("Incorrect previous node value, got: %d, expected: %d.", node.Previous().Value(), prev.Value())
				}
			}

			if next != nil {
				if next.Value() != node.Next().Value() {
					t.Errorf("Incorrect next node value, got: %d, expected: %d.", node.Next().Value(), next.Value())
				}
			}

			if node.Next() != nil {
				next = node.Next()
				next = next.Next()
			} else {
				next = nil
			}

			prev = node
		}
	}
}
