package shared

// Indexer interface provides an access
// to collection elements by non-negative index.
type Indexer interface {
	// ElementAt returns collection element by its non-negative index.
	ElementAt(index int) interface{}
	// SetElementAt assigns new value to the collection element by its non-negative index.
	SetElementAt(index int, value interface{})
}
