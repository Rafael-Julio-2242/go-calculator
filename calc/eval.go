package calculator

import (
	"errors"
	"fmt"
	"strconv"
)

func Eval(n *Node) (float64, error) {

	if n.Left == nil && n.Right == nil { // Caso nó folha
		number, errConvLeaf := strconv.ParseFloat(n.Value, 64)
		if errConvLeaf != nil {
			return 0, errors.New("an operation Cannot be a leaf")
		}

		return number, nil
	}

	leftValue, err := Eval(n.Left)

	if err != nil {
		return 0, err
	}

	rightValue, err := Eval(n.Right)

	if err != nil {
		return 0, err
	}

	switch n.Value {
	case "+":
		return leftValue + rightValue, nil
	case "-":
		return leftValue - rightValue, nil
	case "*":
		return leftValue * rightValue, nil
	case "/":
		if rightValue == 0 {
			return 0, errors.New("zero division is not supported")
		}
		return leftValue / rightValue, nil

	default:
		numberVal, errConvNumberVal := strconv.ParseFloat(n.Value, 64)

		if errConvNumberVal != nil {
			return 0, fmt.Errorf("invalid operator: %q", n.Value)
		}

		return numberVal, nil
	}
}
