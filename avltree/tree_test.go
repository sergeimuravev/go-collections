package avltree_test

import (
	"testing"

	. "github.com/sergeimuravev/go-collections/avltree"
	"github.com/sergeimuravev/go-collections/shared"
)

var IntegerComparer shared.Comparer = func(a, b interface{}) int {
	if a.(int) < b.(int) {
		return -1
	}

	if a.(int) > b.(int) {
		return 1
	}

	return 0
}

func TestAdd(t *testing.T) {
	samples := []struct {
		values   []int
		expected []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4}, []int{2, 1, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 2, 1, 3, 5, 6}},
	}

	for _, sample := range samples {
		var tree = New(IntegerComparer)
		for _, v := range sample.values {
			tree.Add(v, v)
		}

		if tree.Count() != len(sample.values) {
			t.Errorf("Incorrect length, got: %d, expected: %d.", tree.Count(), len(sample.values))
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

func TestRemove(t *testing.T) {
	samples := []struct {
		values   []int
		remove   int
		expected []int
	}{
		{[]int{1}, 0, []int{1}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 2, []int{1}},
		{[]int{1, 2, 3}, 3, []int{2, 1}},
		{[]int{1, 2, 3, 4}, 2, []int{3, 1, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{2, 1, 5, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 4, []int{5, 2, 1, 3, 6}},
	}

	for _, sample := range samples {
		var tree = New(IntegerComparer)
		for _, v := range sample.values {
			tree.Add(v, v)
		}

		count := len(sample.values)
		removed := tree.Remove(sample.remove)
		if removed {
			count--
		}

		if tree.Count() != count {
			t.Errorf("Incorrect length, got: %d, expected: %d.", tree.Count(), count)
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

func TestSearch(t *testing.T) {
	samples := []struct {
		values   []int
		value    int
		expected interface{}
	}{
		{[]int{1}, 0, nil},
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 2},
		{[]int{1, 2, 2}, 2, 2},
		{[]int{1, 2, 3}, 3, 3},
		{[]int{1, 2, 3}, 0, nil},
		{[]int{1, 2, 3, 4}, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 5, 5},
		{[]int{1, 2, 3, 4, 5, 6}, 3, 3},
	}

	for _, sample := range samples {
		var tree = New(IntegerComparer)
		for _, v := range sample.values {
			tree.Add(v, v)
		}

		var found = tree.Get(sample.value)
		if found != sample.expected {
			t.Errorf("Incorrect key, got: %d, expected: %d.", found, sample.expected)
		}
	}
}

func TestClear(t *testing.T) {
	var tree = New(nil)
	values := []int{1, 2, 3}
	for _, v := range values {
		tree.Add(v, v)
	}

	if tree.Count() != len(values) {
		t.Errorf("Incorrect length, got: %d, expected: %d.", tree.Count(), len(values))
	}

	tree.Clear()

	if tree.Count() != 0 {
		t.Errorf("Incorrect length, got: %d, expected: %d.", tree.Count(), 0)
	}
}
