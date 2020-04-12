package linkedlist_test

import (
	. "../linkedlist"
	"fmt"
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

		it := sample.list.GetEnumerator()
		for i := 0; i < length; i++ {
			if !it.MoveNext() {
				break
			}

			node := it.Current().(Node)
			current := node.Value()
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect node value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestFirst(t *testing.T) {
	samples := []struct {
		list  LinkedList
		value interface{}
	}{
		{New(1, 2, 3, 4, 5), 5},
		{New(5, 4, 3, 2, 1), 1},
		{LinkedList{}, 0},
	}

	for _, sample := range samples {
		var current, expected interface{} = 0, sample.value

		firstNode := sample.list.First()
		if firstNode != nil {
			current = firstNode.Value()
		}

		if current != expected {
			t.Errorf("Incorrect node value, got: %d, expected: %d.", current, expected)
		}
	}
}

func TestAdd(t *testing.T) {
	samples := []struct {
		list   LinkedList
		values []int
	}{
		{New(1, 2, 3, 4, 5), []int{0, 5, 4, 3, 2, 1}},
		{New(5, 4, 3, 2, 1), []int{0, 1, 2, 3, 4, 5}},
		{New(), []int{0}},
	}

	for _, sample := range samples {
		count := sample.list.Count()
		sample.list.Add(0)

		if sample.list.Count()-count != 1 {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), count)
		}

		it := sample.list.GetEnumerator()
		for i := 0; i < len(sample.values); i++ {
			if !it.MoveNext() {
				break
			}

			node := it.Current().(Node)
			current := node.Value()
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect node value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestRemove(t *testing.T) {
	samples := []struct {
		list   LinkedList
		value  int
		values []int
		result bool
	}{
		{New(1, 2, 3), 1, []int{3, 2}, true},
		{New(1, 2, 3), 2, []int{3, 1}, true},
		{New(1, 2, 3), 0, nil, false},
		{New(1), 0, nil, false},
		{LinkedList{}, 0, nil, false},
	}

	for _, sample := range samples {
		count := sample.list.Count()
		node := sample.list.Find(sample.value)
		result := sample.list.Remove(node)
		if result != sample.result {
			t.Errorf("Incorrect result, got: %v, expected: %v.", result, sample.result)
		}

		if result == true {
			if count-sample.list.Count() != 1 {
				t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), count)
			}

		}

		it := sample.list.GetEnumerator()
		for _, expected := range sample.values {
			it.MoveNext()
			node := it.Current().(Node)
			if node.Value() != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", node.Value(), expected)
			}
		}
	}
}

func TestFind(t *testing.T) {
	samples := []struct {
		list   LinkedList
		value  int
		exists bool
	}{
		{New(1, 2, 3, 4, 5), 1, true},
		{New(1, 2, 3, 4, 5), 2, true},
		{New(1, 2, 3, 4, 5), 3, true},
		{New(1, 2, 3, 4, 5), 4, true},
		{New(1, 2, 3, 4, 5), 5, true},
		{New(1, 2, 3, 4, 5), 0, false},
	}

	for _, sample := range samples {
		node := sample.list.Find(sample.value)
		if sample.exists != (node != nil) {
			t.Errorf("Incorrect result, got: %v, expected: %v.", node != nil, sample.exists)
		}

		if sample.exists && node.Value() != sample.value {
			t.Errorf("Incorrect node value, got: %d, expected: %d.", node.Value(), sample.value)
		}
	}
}

func TestClear(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	list.Clear()

	if list.Count() != 0 {
		t.Errorf("Incorrect count, got: %d, expected: %d.", list.Count(), 0)
	}

	if list.First() != nil {
		t.Error("First node should be nil.")
	}
}

func TestListStringer(t *testing.T) {
	list := New(1, 2, 3, 4, 5)
	var stringer fmt.Stringer = &list
	value := stringer.String()
	expected := fmt.Sprintf("Linked list. Count=%d.", list.Count())
	if value != expected {
		t.Errorf("Incorrect value, got: %s, expected: %s.", value, expected)
	}
}
