package models

type Node struct {
	failure   *Node
	transport map[byte]*Node
	children  map[byte]*Node
}
