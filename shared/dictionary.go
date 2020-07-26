package shared

// Dictionary interface represents an associated array which values are
// retrievable by key
type Dictionary interface {
	// Add adds an element with comparable key to collection.
	Add(key, value interface{})

	// Get returns value by key.
	Get(key interface{}) interface{}

	// Remove removes element from collection.
	// Returns true if element with provided key exists
	// in collection, otherwise returns false.
	Remove(key interface{}) bool

	// ContainsKey returns true if collection contains an element with provided key.
	ContainsKey(key interface{}) bool
}
