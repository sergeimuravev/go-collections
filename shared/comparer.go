package shared

// Comparer is used to compare two values.
// Returns 0 if both values are equal, positive int if the left value exceeds right, negative int otherwise.
type Comparer func(interface{}, interface{}) int
