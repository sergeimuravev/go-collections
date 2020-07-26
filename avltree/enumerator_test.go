package avltree_test

import (
	"testing"

	. "github.com/sergeimuravev/go-collections/avltree"
)

func TestMoveNext(t *testing.T) {
	samples := []struct {
		values   []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{4, 2, 1, 3, 6, 5, 7, 8}},
	}

	for _, sample := range samples {
		var tree = New(IntegerComparer)
		for _, v := range sample.values {
			tree.Add(v, v)
		}

		i := 0
		it := tree.GetEnumerator()
		for it.MoveNext() {
			node := *it.Current().(*Node)
			expected := sample.expected[i]
			if node.Key != expected {
				t.Errorf("Incorrect key, got: %d, expected: %d.", node.Key, expected)
			}
			i++
		}
	}
}

func TestReset(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}
	var tree = New(IntegerComparer)
	for _, v := range values {
		tree.Add(v, v)
	}

	it := tree.GetEnumerator()
	for it.MoveNext() {

	}

	node := *it.Current().(*Node)
	expected := values[len(values)-1]
	if node.Key != expected {
		t.Errorf("Incorrect key, got: %d, expected: %d.", node.Key, expected)
	}

	it.Reset()
	it.MoveNext()

	node = *it.Current().(*Node)
	expected = 2 // See test above
	if node.Key != expected {
		t.Errorf("Incorrect key, got: %d, expected: %d.", node.Key, expected)
	}
}
