package doublylinkedlist

// Node type specifies element type in the doubly linked list.
type Node struct {
	value    interface{}
	next     *Node
	previous *Node
}

// NewNode creates linked list node.
func NewNode(value interface{}, next *Node, previous *Node) Node {
	return Node{value, next, previous}
}

// Value returns linked list node value.
func (node *Node) Value() interface{} {
	return node.value
}

// SetValue assigns linked list node value.
func (node *Node) SetValue(newValue interface{}) {
	node.value = newValue
}

// Next returns pointer to the node next to current one.
func (node *Node) Next() *Node {
	return node.next
}

// SetNext assigns node next to current one.
func (node *Node) SetNext(newNode *Node) {
	node.next = newNode
}

// Previous returns pointer to the node previous to current one.
func (node *Node) Previous() *Node {
	return node.previous
}

// SetPrevious assigns node previous to current one.
func (node *Node) SetPrevious(newNode *Node) {
	node.previous = newNode
}
