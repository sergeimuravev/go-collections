package heap_test

import (
	"fmt"
	"github.com/sergeimuravev/go-collections/heap"
)

func Example() {
	// Create heap with some values
	collection := append(make([]interface{}, 0), 2, 1, 5, 3, 4)
	comparer := func(x interface{}, y interface{}) int {
		return x.(int) - y.(int)
	}

	heap := heap.New(collection, comparer)

	// Print internal values
	it := heap.GetEnumerator()
	for it.MoveNext() {
		fmt.Println(it.Current())
	}

	// Peek max value
	fmt.Println(heap.Peek())

	// Push some data
	heap.Push(0)
	heap.Push(10)
	heap.Push(10)
	// Print heap count
	fmt.Println(heap.Count())

	// Pop sorted values
	for heap.Count() > 0 {
		fmt.Println(heap.Pop())
	}

	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
	// 5
	// 8
	// 10
	// 10
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
