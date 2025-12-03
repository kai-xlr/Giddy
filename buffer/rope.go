package buffer

import "fmt"

// Rope is an immutable wrapper around the tree root.
type Rope struct {
	Root *Node
}

// NewRope creates a rope from a standard string.
// Strategy: For now, just create a single Leaf node.
// Later optimization: If string is huge, recursively split it.
func NewRope(initialText string) *Rope {
	panic("implement me")
}

// Bytes flattens the rope back into a byte slice.
// Strategy: Traverse the tree (In-Order Traversal) and append leaf values.
func (r *Rope) Bytes() []byte {
	panic("implement me")
}

// String returns the string representation (useful for debugging).
func (r *Rope) String() string {
	return string(r.Bytes())
}

// Insert adds text at a specific index, returning a NEW Rope (immutability).
// Strategy:
// 1. Split the tree at `pos` into left and right.
// 2. Create a new node for the text.
// 3. Concat(left, new_node) -> temp.
// 4. Concat(temp, right) -> new_root.
func (r *Rope) Insert(pos int, text string) *Rope {
	panic("implement me")
}

// Delete removes text between start and end, returning a NEW Rope.
// Strategy:
// 1. Split at `start` -> (keep_left, temp_right).
// 2. Split `temp_right` at (end - start) -> (trash, keep_right).
// 3. Concat(keep_left, keep_right).
func (r *Rope) Delete(start, end int) *Rope {
	panic("implement me")
}

// --- Internal Helpers (The "Hard" Logic) ---

// concat joins two nodes into a new root.
// Hint: This involves creating a new Internal Node where Left=left and Right=right.
// You must update the new root's Weight and Newlines correctly.
func concat(left, right *Node) *Node {
	panic("implement me")
}

// split divides a node at index i into two new nodes.
// This is usually a recursive function.
func split(n *Node, i int) (*Node, *Node) {
	panic("implement me")
}
