package list_test

import (
	. "github.com/sergeimuravev/go-collections/list"
	"testing"
)

func TestNew(t *testing.T) {
	samples := []struct {
		list   List
		values []int
	}{
		{New(1, 2, 3, 4, 5), []int{1, 2, 3, 4, 5}},
		{New(5, 4, 3, 2, 1), []int{5, 4, 3, 2, 1}},
		{New(), []int{}},
		{New(make([]interface{}, 0)), []int{}},
		{New(append(make([]interface{}, 0), 1, 2, 3, 4, 5)), []int{1, 2, 3, 4, 5}},
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

			current := it.Current()
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestIndexer(t *testing.T) {
	samples := []struct {
		list  List
		index int
		value int
	}{
		{New(1, 2, 3, 4, 5), 0, 100},
		{New(1, 2, 3, 4, 5), 1, 200},
		{New(1, 2, 3, 4, 5), 2, 300},
		{New(1, 2, 3, 4, 5), 3, 400},
		{New(1, 2, 3, 4, 5), 4, 500},
	}

	for _, sample := range samples {
		sample.list.SetElementAt(sample.index, sample.value)
		value := sample.list.ElementAt(sample.index).(int)
		if value != sample.value {
			t.Errorf("Incorrect updated value, got: %d, expected: %d.", value, sample.value)
		}
	}
}

func TestAdd(t *testing.T) {
	samples := []struct {
		list   List
		value  int
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, []int{1, 2, 3, 4, 5, 0}},
		{New(5, 4, 3, 2, 1), 10, []int{5, 4, 3, 2, 1, 10}},
		{New(), 0, []int{0}},
	}

	for _, sample := range samples {
		sample.list.Add(sample.value)
		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestAddAll(t *testing.T) {
	samples := []struct {
		list   List
		set    []interface{}
		values []int
	}{
		{New(1, 2, 3, 4, 5), append(make([]interface{}, 0), 1, 2, 3), []int{1, 2, 3, 4, 5, 1, 2, 3}},
		{New(5, 4, 3, 2, 1), append(make([]interface{}, 0), 3, 2, 1), []int{5, 4, 3, 2, 1, 3, 2, 1}},
		{New(), append(make([]interface{}, 0), 0), []int{0}},
	}

	for _, sample := range samples {
		sample.list.AddAll(sample.set)

		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestInsertAt(t *testing.T) {
	samples := []struct {
		list   List
		index  int
		value  interface{}
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, 0, []int{0, 1, 2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 1, 0, []int{1, 0, 2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 2, 0, []int{1, 2, 0, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 3, 0, []int{1, 2, 3, 0, 4, 5}},
		{New(1, 2, 3, 4, 5), 4, 0, []int{1, 2, 3, 4, 0, 5}},
		{New(), 0, 0, []int{0}},
	}

	for _, sample := range samples {
		sample.list.InsertAt(sample.index, sample.value)
		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestInsertAll(t *testing.T) {
	samples := []struct {
		list   List
		index  int
		insert []interface{}
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, append(make([]interface{}, 0), 100, 200, 300), []int{100, 200, 300, 1, 2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 1, append(make([]interface{}, 0), 100, 200, 300), []int{1, 100, 200, 300, 2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 2, append(make([]interface{}, 0), 100, 200, 300), []int{1, 2, 100, 200, 300, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 3, append(make([]interface{}, 0), 100, 200, 300), []int{1, 2, 3, 100, 200, 300, 4, 5}},
		{New(1, 2, 3, 4, 5), 4, append(make([]interface{}, 0), 100, 200, 300), []int{1, 2, 3, 4, 100, 200, 300, 5}},
		{New(1, 2, 3, 4, 5), 5, append(make([]interface{}, 0), 100, 200, 300), []int{1, 2, 3, 4, 5, 100, 200, 300}},
	}

	for _, sample := range samples {
		sample.list.InsertAll(sample.index, sample.insert)
		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestIndexOf(t *testing.T) {
	samples := []struct {
		list  List
		value interface{}
		index int
	}{
		{New(1, 2, 3, 4, 5), 0, -1},
		{New(1, 2, 3, 4, 5), 1, 0},
		{New(1, 2, 3, 4, 5), 2, 1},
		{New(1, 2, 3, 4, 5), 3, 2},
		{New(1, 2, 3, 4, 5), 4, 3},
		{New(1, 2, 3, 4, 5), 5, 4},
	}

	for _, sample := range samples {
		index := sample.list.IndexOf(sample.value)
		if index != sample.index {
			t.Errorf("Incorrect index of element, got: %d, expected: %d.", index, sample.index)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	samples := []struct {
		list  List
		value interface{}
		index int
	}{
		{New(1, 1, 1, 1, 1), 0, -1},
		{New(1, 1, 1, 1, 1), 1, 4},
		{New(1, 2, 3, 3, 4, 5), 3, 3},
		{New(1, 2, 3, 3, 4, 5), 2, 1},
	}

	for _, sample := range samples {
		index := sample.list.LastIndexOf(sample.value)
		if index != sample.index {
			t.Errorf("Incorrect index of element, got: %d, expected: %d.", index, sample.index)
		}
	}
}

func TestFindAll(t *testing.T) {
	samples := []struct {
		list   List
		value  interface{}
		values []interface{}
	}{
		{New(1, 2, 2, 3, 3), 0, make([]interface{}, 0)},
		{New(1, 2, 2, 3, 3), 1, append(make([]interface{}, 0), 1)},
		{New(1, 2, 2, 3, 3), 2, append(make([]interface{}, 0), 2, 2)},
		{New(1, 2, 2, 3, 3), 3, append(make([]interface{}, 0), 3, 3)},
	}

	for _, sample := range samples {
		values := sample.list.FindAll(func(v interface{}) bool {
			return v == sample.value
		})

		for i := range sample.values {
			current := values[i]
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestRemoveAt(t *testing.T) {
	samples := []struct {
		list   List
		index  int
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, []int{2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 1, []int{1, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 2, []int{1, 2, 4, 5}},
		{New(1, 2, 3, 4, 5), 3, []int{1, 2, 3, 5}},
		{New(1, 2, 3, 4, 5), 4, []int{1, 2, 3, 4}},
	}

	for _, sample := range samples {
		sample.list.RemoveAt(sample.index)
		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestRemoveAll(t *testing.T) {
	samples := []struct {
		list   List
		remove int
		values []int
	}{
		{New(1, 2, 3, 4, 5), 0, []int{1, 2, 3, 4, 5}},
		{New(1, 2, 3, 4, 5), 1, []int{2, 3, 4, 5}},
		{New(1, 1, 3, 4, 5), 1, []int{3, 4, 5}},
		{New(1, 1, 3, 4, 1), 1, []int{3, 4}},
		{New(1, 1, 1, 1, 1), 1, []int{}},
	}

	for _, sample := range samples {
		sample.list.RemoveAll(sample.remove)
		for i := range sample.values {
			current := sample.list.ElementAt(i)
			expected := sample.values[i]
			if current != expected {
				t.Errorf("Incorrect value, got: %d, expected: %d.", current, expected)
			}
		}
	}
}

func TestTrimExcess(t *testing.T) {
	list := New()
	for i := 0; i < 1024; i++ {
		list.Add(i % 2)
	}

	before := list.Capacity()
	list.RemoveAll(0)
	after := list.Capacity()

	if after >= before {
		t.Errorf("Incorrect capacity, got: %d, before: %d.", after, before)
	}
}
