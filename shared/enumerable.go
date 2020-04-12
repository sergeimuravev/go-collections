package shared

// Enumerable interface provides a method to retrieve an iterator for collection.
type Enumerable interface {
	GetEnumerator() Enumerator
}
