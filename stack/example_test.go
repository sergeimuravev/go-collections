package stack_test

import (
	"github.com/sergeimuravev/go-collections/linkedlist"
	"github.com/sergeimuravev/go-collections/stack"
	"fmt"
)

func Example() {
	stack := stack.New(1, 2, 3)

	// Push
	stack.Push(4)
	stack.Push(5)

	// Count
	count := stack.Count()
	fmt.Println(count)

	// Peek
	next := stack.Peek()
	fmt.Println(next)

	// Enumerate
	it := stack.GetEnumerator()
	for it.MoveNext() {
		node := it.Current().(*linkedlist.Node)
		fmt.Println(node.Value())
	}

	// Pop
	for stack.Count() > 0 {
		next = stack.Pop()
		fmt.Println(next)
	}

	// Output:
	// 5
	// 5
	// 5
	// 4
	// 3
	// 2
	// 1
	// 5
	// 4
	// 3
	// 2
	// 1
}
