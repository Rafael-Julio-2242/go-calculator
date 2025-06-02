package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

func (N *NodeStack) Length() int {
	return len(N.nodes)
}

func (N *NodeStack) IsEmpty() bool {
	return len(N.nodes) == 0
}

func MountTree(postfixExpression []string) (*Node, error) {

	var stack NodeStack

	operatorsRange := "+-/*"

	for _, ex := range postfixExpression {
		number := true

		// Posso tentar fazer por conversão
		_, errNumConvert := strconv.ParseFloat(ex, 64)

		if errNumConvert != nil {
			number = false
		}

		if number {
			node := Node{
				Value: ex,
				Left:  nil,
				Right: nil,
			}

			stack.Push(&node)
			continue
		}

		if !number && strings.Contains(operatorsRange, ex) {

			top1, errPop1 := stack.Pop()
			top2, errPop2 := stack.Pop()

			if errPop1 != nil || errPop2 != nil {
				return &Node{Value: "", Left: nil, Right: nil}, errors.New("error on poping values")
			}

			operatorNode := Node{
				Value: ex,
				Right: top1,
				Left:  top2,
			}

			stack.Push(&operatorNode)
		} else if !strings.Contains(operatorsRange, ex) {
			errorMsg := fmt.Sprintf("invalid operato: %s", ex)
			return &Node{Value: "", Left: nil, Right: nil}, errors.New(errorMsg)
		}

	}

	if stack.Length() > 1 {
		return &Node{Value: "", Left: nil, Right: nil}, errors.New("expression error: expression tree has more than one root")
	}

	rootNode, errPopRoot := stack.Pop()

	if errPopRoot != nil {
		return &Node{Value: "", Left: nil, Right: nil}, errors.New("error poping root")
	}

	return rootNode, nil
}

func PrintTree(node *Node, indent string) {
	if node == nil {
		return
	}

	fmt.Println(indent + node.Value)

	PrintTree(node.Left, indent+" ")
	PrintTree(node.Right, indent+" ")
}
