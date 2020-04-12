package shared

// Indexer interface provides an access
// to collection elements by non-negative index.
type Indexer interface {
	// Returns collection element by its non-negative index.
	ElementAt(index int) interface{}
}
