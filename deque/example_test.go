package deque_test

import (
	"../deque"
	"../doublylinkedlist"
	"fmt"
)

func Example() {
	deque := deque.New(2, 3, 4)

	// Push front
	deque.PushFront(1)

	// Push back
	deque.PushBack(5)

	// Count
	count := deque.Count()
	fmt.Println(count)

	// Peek front
	next := deque.Front()
	fmt.Println(next)

	// Peek back
	next = deque.Back()
	fmt.Println(next)

	// Enumerate
	it := deque.GetEnumerator()
	for it.MoveNext() {
		node := it.Current().(*doublylinkedlist.Node)
		fmt.Println(node.Value())
	}

	// Pop front
	next = deque.PopFront()
	fmt.Println(next)

	// Pop back
	next = deque.PopBack()
	fmt.Println(next)

	// Output:
	// 5
	// 1
	// 5
	// 5
	// 4
	// 3
	// 2
	// 1
	// 1
	// 5
}
