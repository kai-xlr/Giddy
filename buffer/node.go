package buffer

// Node represents a node in the Rope binary tree.
// If Left and Right are nil, this is a Leaf node containing raw text in Value.
// If Left and/or Right are not nil, this is an Internal node.
type Node struct {
	Left  *Node
	Right *Node

	// Value holds the text for Leaf nodes only.
	// For internal nodes, this should be nil or empty.
	Value []byte

	// Weight is the character count of the LEFT subtree (or this node if it's a leaf).
	// This is critical for O(log N) indexing.
	Weight int

	// Newlines tracks the number of '\n' characters in this node's subtree.
	// This is critical for line counting.
	Newlines int
}

// NewLeaf creates a simple leaf node containing text.
// Hint: You need to calculate the Weight (len(text)) and count the newlines here.
func NewLeaf(text []byte) *Node {
	panic("implement me")
}

// TotalWeight returns the total length of text in this subtree.
// Caution: For internal nodes, this is usually Weight + Right.TotalWeight()
func (n *Node) TotalWeight() int {
	panic("implement me")
}
