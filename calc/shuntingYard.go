package calculator

import (
	"errors"
	"strings"
)

func ShuntingYard(expression string) ([]string, error) {

	var operatorStack []string
	var operandStack []string

	numericRange := "123456789"

	for _, s := range expression {

		if strings.Contains(numericRange, string(s)) {
			operandStack = append(operandStack, string(s))
			continue
		}

		currentPrecedence, currentAssociativity, err := getPrecedenceInfo(s)

		lastIndex := len(operandStack) - 1

		lastPrecedence, lastAssociativity, _ := getPrecedenceInfo(rune(operandStack[lastIndex][0]))

		if err != nil {
			return []string{""}, err
		}

		if currentPrecedence > lastPrecedence || (currentPrecedence == lastPrecedence && currentAssociativity == lastAssociativity && currentAssociativity == "right") {
			operatorStack = append(operatorStack, string(s))
			continue
		}

		if currentPrecedence < lastPrecedence || (currentPrecedence == lastPrecedence && currentAssociativity == lastAssociativity && currentAssociativity == "left") {
			value := operatorStack[len(operatorStack)-1]
			operatorStack[len(operatorStack)-1] = ""
			operatorStack = operatorStack[:len(operatorStack)-1]
			operandStack = append(operandStack, value)
			continue
		}

	}

	return operandStack, nil
}

func getPrecedenceInfo(r rune) (int, string, error) {
	switch r {
	case '^':
		return 4, "right", nil
	case '*':
		return 3, "left", nil
	case '/':
		return 3, "left", nil
	case '+':
		return 2, "left", nil
	case '-':
		return 2, "left", nil
	default:
		return 0, "", errors.New("invalid operator informed")
	}
}
