package doublylinkedlist_test

import (
	. "github.com/sergeimuravev/go-collections/doublylinkedlist"
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
		{New(make([]interface{}, 0)), []int{}},
		{New(append(make([]interface{}, 0), 1, 2, 3, 4, 5)), []int{5, 4, 3, 2, 1}},
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

func TestAddFirst(t *testing.T) {
	samples := []struct {
		list   LinkedList
		value  int
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, []int{0, 5, 4, 3, 2, 1}},
		{New(5, 4, 3, 2, 1), 0, []int{0, 1, 2, 3, 4, 5}},
		{New(), 0, []int{0}},
	}

	for _, sample := range samples {
		sample.list.AddFirst(sample.value)

		length := len(sample.values)
		if sample.list.Count() != length {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), length)
		}

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
		}
	}
}

func TestAddLast(t *testing.T) {
	samples := []struct {
		list   LinkedList
		value  int
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, []int{5, 4, 3, 2, 1, 0}},
		{New(5, 4, 3, 2, 1), 0, []int{1, 2, 3, 4, 5, 0}},
		{New(), 0, []int{0}},
	}

	for _, sample := range samples {
		sample.list.AddLast(sample.value)

		length := len(sample.values)
		if sample.list.Count() != length {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), length)
		}

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
		}
	}
}

func TestAddBefore(t *testing.T) {
	samples := []struct {
		list        LinkedList
		valueBefore int
		value       int
		values      []int
	}{
		{New(1, 2, 3, 4, 5), 1, 0, []int{5, 4, 3, 2, 0, 1}},
		{New(1, 2, 3, 4, 5), 2, 0, []int{5, 4, 3, 0, 2, 1}},
		{New(1, 2, 3, 4, 5), 3, 0, []int{5, 4, 0, 3, 2, 1}},
		{New(1, 2, 3, 4, 5), 4, 0, []int{5, 0, 4, 3, 2, 1}},
		{New(1, 2, 3, 4, 5), 5, 0, []int{0, 5, 4, 3, 2, 1}},
		{New(5, 4, 3, 2, 1), 1, 0, []int{0, 1, 2, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), 2, 0, []int{1, 0, 2, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), 3, 0, []int{1, 2, 0, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), 4, 0, []int{1, 2, 3, 0, 4, 5}},
		{New(5, 4, 3, 2, 1), 5, 0, []int{1, 2, 3, 4, 0, 5}},
		{New(), 1, 0, []int{0}},
	}

	for _, sample := range samples {
		node := sample.list.Find(nil, nil, sample.valueBefore)
		sample.list.AddBefore(node, sample.value)

		length := len(sample.values)
		if sample.list.Count() != length {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), length)
		}

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
		}
	}
}

func TestAddAfter(t *testing.T) {
	samples := []struct {
		list        LinkedList
		valueBefore int
		value       int
		values      []int
	}{
		{New(1, 2, 3, 4, 5), 1, 0, []int{5, 4, 3, 2, 1, 0}},
		{New(1, 2, 3, 4, 5), 2, 0, []int{5, 4, 3, 2, 0, 1}},
		{New(1, 2, 3, 4, 5), 3, 0, []int{5, 4, 3, 0, 2, 1}},
		{New(1, 2, 3, 4, 5), 4, 0, []int{5, 4, 0, 3, 2, 1}},
		{New(1, 2, 3, 4, 5), 5, 0, []int{5, 0, 4, 3, 2, 1}},
		{New(5, 4, 3, 2, 1), 1, 0, []int{1, 0, 2, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), 2, 0, []int{1, 2, 0, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), 3, 0, []int{1, 2, 3, 0, 4, 5}},
		{New(5, 4, 3, 2, 1), 4, 0, []int{1, 2, 3, 4, 0, 5}},
		{New(5, 4, 3, 2, 1), 5, 0, []int{1, 2, 3, 4, 5, 0}},
		{New(), 1, 0, []int{0}},
	}

	for _, sample := range samples {
		node := sample.list.Find(nil, nil, sample.valueBefore)
		sample.list.AddAfter(node, sample.value)

		length := len(sample.values)
		if sample.list.Count() != length {
			t.Errorf("Incorrect count, got: %d, expected: %d.", sample.list.Count(), length)
		}

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
		}
	}
}

func TestFind(t *testing.T) {
	samples := []struct {
		list   LinkedList
		from   interface{}
		to     interface{}
		value  int
		exists bool
	}{
		{New(1, 2, 3, 4, 5), nil, nil, 3, true},
		{New(1, 2, 3, 4, 5), 5, nil, 3, true},
		{New(1, 2, 3, 4, 5), nil, 1, 3, true},
		{New(1, 2, 3, 4, 5), nil, nil, 0, false},
		{New(1, 2, 3, 4, 5), 4, 2, 3, true},
		{New(1, 2, 3, 4, 5), 4, 2, 0, false},
		{New(1, 2, 3, 4, 5), 4, 3, 3, true},
		{New(1, 2, 3, 4, 5), 3, 2, 3, true},
		{New(1, 2, 3, 4, 5), 3, 3, 3, true},
		{New(1, 2, 3, 4, 5), nil, 3, 3, true},
		{New(1, 2, 3, 4, 5), 3, nil, 3, true},

		{New(1, 2, 3, 4), nil, nil, 3, true},
		{New(1, 2, 3, 4), 4, nil, 3, true},
		{New(1, 2, 3, 4), nil, 1, 3, true},
		{New(1, 2, 3, 4), nil, nil, 0, false},
		{New(1, 2, 3, 4), 4, 2, 3, true},
		{New(1, 2, 3, 4), 4, 2, 0, false},
		{New(1, 2, 3, 4), 4, 3, 3, true},
		{New(1, 2, 3, 4), 3, 2, 3, true},
		{New(1, 2, 3, 4), 3, 3, 3, true},
		{New(1, 2, 3, 4), nil, 3, 3, true},
		{New(1, 2, 3, 4), 3, nil, 3, true},

		{New(), nil, nil, 0, false},
		{New(0), nil, nil, 0, true},
	}

	for _, sample := range samples {
		left := sample.list.Find(nil, nil, sample.from)
		right := sample.list.Find(nil, nil, sample.to)
		node := sample.list.Find(left, right, sample.value)
		if sample.exists != (node != nil) {
			t.Errorf("Incorrect result, got: %v, expected: %v.", node != nil, sample.exists)
		}

		if sample.exists && node.Value() != sample.value {
			t.Errorf("Incorrect node value, got: %d, expected: %d.", node.Value(), sample.value)
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
		{New(1, 2, 3, 4, 5), 1, []int{5, 4, 3, 2}, true},
		{New(1, 2, 3, 4, 5), 2, []int{5, 4, 3, 1}, true},
		{New(1, 2, 3, 4, 5), 3, []int{5, 4, 2, 1}, true},
		{New(1, 2, 3, 4, 5), 4, []int{5, 3, 2, 1}, true},
		{New(1, 2, 3, 4, 5), 5, []int{4, 3, 2, 1}, true},
		{New(1, 2, 3, 4, 5), 0, []int{5, 4, 3, 2, 1}, false},
		{New(1), 0, []int{1}, false},
		{LinkedList{}, 0, []int{}, false},
	}

	for _, sample := range samples {
		count := sample.list.Count()
		node := sample.list.Find(nil, nil, sample.value)
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
			node := it.Current().(*Node)
			if node.Value() != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", node.Value(), expected)
			}
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

	if list.Last() != nil {
		t.Error("Last node should be nil.")
	}
}
