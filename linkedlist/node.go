package linkedlist

// Node type specifies element type in the linked list.
type Node struct {
	value interface{}
	next  *Node
}

// NewNode creates linked list node.
func NewNode(value interface{}, next *Node) Node {
	return Node{value, next}
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
