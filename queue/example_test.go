package queue_test

import (
	"fmt"

	"github.com/sergeimuravev/go-collections/linkedlist"
	"github.com/sergeimuravev/go-collections/queue"
)

func Example() {
	queue := queue.New(1, 2, 3)

	// Enqueue
	queue.Enqueue(4)
	queue.Enqueue(5)

	// Count
	count := queue.Count()
	fmt.Println(count)

	// Peek
	next := queue.Peek()
	fmt.Println(next)

	// Enumerate
	it := queue.GetEnumerator()
	for it.MoveNext() {
		node := it.Current().(*linkedlist.Node)
		fmt.Println(node.Value())
	}

	// Dequeue
	for queue.Count() > 0 {
		next = queue.Dequeue()
		fmt.Println(next)
	}

	// Output:
	// 5
	// 1
	// 1
	// 2
	// 3
	// 4
	// 5
	// 1
	// 2
	// 3
	// 4
	// 5
}
