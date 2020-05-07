package linkedlist_test

import (
	. "github.com/sergeimuravev/go-collections/linkedlist"
	"testing"
)

func TestNewNode(t *testing.T) {
	last := NewNode(1, nil)
	first := NewNode(0, &last)

	if *first.Next() != last {
		t.Error("First node should point to the last one.")
	}

	if first.Value() != 0 {
		t.Errorf("Incorrect first node value, got: %d, expected: %d.", first.Value(), 0)
	}

	if last.Next() != nil {
		t.Error("Last node should point to nil.")
	}

	if last.Value() != 1 {
		t.Errorf("Incorrect last node value, got: %d, expected: %d.", last.Value(), 1)
	}
}

func TestSetters(t *testing.T) {
	next := NewNode(1, nil)
	node := NewNode(0, nil)

	node.SetValue(2)
	node.SetNext(&next)

	if node.Value() != 2 {
		t.Errorf("Incorrect value, got: %d, expected: %d.", node.Value(), 2)
	}

	if node.Next() != &next {
		t.Error("Incorrect next node reference.")
	}
}
