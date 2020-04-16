package doublylinkedlist_test

import (
	. "../doublylinkedlist"
	"testing"
)

func TestNewNode(t *testing.T) {
	var first Node
	second := NewNode(2, nil, &first)
	first = NewNode(1, &second, nil)

	if first.Next() != &second {
		t.Error("First node should point to the second one.")
	}

	if first.Previous() != nil {
		t.Error("First node should not have previous one.")
	}

	if first.Value() != 1 {
		t.Errorf("Incorrect first node value, got: %d, expected: %d.", first.Value(), 1)
	}

	if second.Next() != nil {
		t.Error("second node should point to nil.")
	}

	if second.Previous() != &first {
		t.Error("second node should point to the first one.")
	}

	if second.Value() != 2 {
		t.Errorf("Incorrect second node value, got: %d, expected: %d.", second.Value(), 2)
	}
}

func TestSetters(t *testing.T) {
	first := Node{}
	second := Node{}

	first.SetValue(1)
	first.SetNext(&second)
	first.SetPrevious(nil)

	second.SetValue(2)
	second.SetNext(nil)
	second.SetPrevious(&first)

	if first.Next() != &second {
		t.Error("First node should point to the second one.")
	}

	if first.Previous() != nil {
		t.Error("First node should not have previous one.")
	}

	if first.Value() != 1 {
		t.Errorf("Incorrect first node value, got: %d, expected: %d.", first.Value(), 0)
	}

	if second.Next() != nil {
		t.Error("second node should point to nil.")
	}

	if second.Previous() != &first {
		t.Error("second node should point to the first one.")
	}

	if second.Value() != 2 {
		t.Errorf("Incorrect second node value, got: %d, expected: %d.", second.Value(), 1)
	}
}
