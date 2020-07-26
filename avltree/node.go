package avltree

// Node is an avl tree node representation
type Node struct {
	left   *Node
	right  *Node
	height uint8
	Key    interface{}
	Value  interface{}
}

func calcHeight(node *Node) uint8 {
	if node == nil {
		return 0
	}

	return node.height
}

// fixHeight restores node's height correct value
func (node *Node) fixHeight() {
	hl := calcHeight(node.left)
	hr := calcHeight(node.right)

	if hl > hr {
		node.height = hl + 1
	} else {
		node.height = hr + 1
	}
}

// balanceFactor calculates node's balance factor (-1, 0 or +1)
func (node *Node) balanceFactor() int8 {
	return int8(calcHeight(node.right) - calcHeight(node.left))
}

// rotateRight performs right turn around node
func (node *Node) rotateRight() *Node {
	q := node.left
	node.left = q.right
	q.right = node
	node.fixHeight()
	q.fixHeight()
	return q
}

// rotateLeft performs left turn around node
func (node *Node) rotateLeft() *Node {
	p := node.right
	node.right = p.left
	p.left = node
	node.fixHeight()
	p.fixHeight()
	return p
}

// balance performs node sub-tree balancing using rotations
func (node *Node) balance() *Node {
	node.fixHeight()

	if node.balanceFactor() > 1 {
		if node.right != nil && node.right.balanceFactor() < 0 {
			node.right = node.right.rotateRight()
		}

		return node.rotateLeft()
	}

	if node.balanceFactor() < -1 {
		if node.left != nil && node.left.balanceFactor() > 0 {
			node.left = node.left.left.rotateLeft()
		}

		return node.rotateRight()
	}

	return node // No re-balancing is required
}

// findMin returns child node with minimal key
func (node *Node) findMin() *Node {
	if node.left != nil {
		return node.left.findMin()
	}

	return node
}

// removeMin deletes node with minimal key
func (node *Node) removeMin() *Node {
	if node.left == nil {
		return node.right
	}

	node.left = node.left.removeMin()
	return node.balance()
}
