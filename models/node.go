package models

type Node struct {
	isEndpoint bool
	failure    *Node
	transport  map[int32]*Node
	children   map[int32]*Node
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) GetChild(b int32, createAnyway bool) *Node {
	if createAnyway && node.children[b] == nil {
		node.children[b] = NewNode()
	}
	return node.children[b]
}
