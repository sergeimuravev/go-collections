package shared

// Enumerator interface provides a method to enumerate collection in forward direction.
type Enumerator interface {
	// Move returns a flag indicating
	// whether more elements can be fetched.
	MoveNext() bool

	// Current returns an element enumerator points to.
	Current() interface{}

	// Reset clears state and resets enumeration to the initial state.
	Reset()
}
