package calculator

import "errors"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}
type NodeStack struct {
	nodes []*Node
}

func (N *NodeStack) Push(n *Node) {
	N.nodes = append(N.nodes, n)
}

func (N *NodeStack) Pop() (*Node, error) {
	if len(N.nodes) == 0 {
		return nil, errors.New("stack is empty")
	}
	lastNode := N.nodes[len(N.nodes)-1]
	N.nodes[len(N.nodes)-1] = nil
	N.nodes = N.nodes[:len(N.nodes)-1]
	return lastNode, nil
}

func (N *NodeStack) Peek() (*Node, error) {
	if len(N.nodes) == 0 {
		return nil, errors.New("stack is empty")
	}

	return N.nodes[len(N.nodes)-1], nil
}

func (N *NodeStack) IsEmpty() bool {
	return len(N.nodes) == 0
}

func MountTree(expression string) {

}
