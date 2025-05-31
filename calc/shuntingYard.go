package calculator

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

func ShuntingYard(expression string) ([]string, error) {

	var operatorStack []string
	var operandStack []string

	numericRange := "123456789"

	for _, s := range expression {

		fmt.Println("[s]: ", string(s))
		fmt.Println()
		fmt.Println("operandStack: ", operandStack)
		fmt.Println("operatorStack: ", operatorStack)

		if strings.Contains(numericRange, string(s)) {
			operandStack = append(operandStack, string(s))
			continue
		}

		currentPrecedence, currentAssociativity, err := getPrecedenceInfo(string(s))

		if err != nil {
			return []string{""}, err
		}

		if len(operatorStack) == 0 {
			operatorStack = append(operatorStack, string(s))
			continue
		}

		lastIndex := len(operatorStack) - 1

		lastPrecedence, lastAssociativity, _ := getPrecedenceInfo(operatorStack[lastIndex])

		if currentPrecedence > lastPrecedence || (currentPrecedence == lastPrecedence && currentAssociativity == lastAssociativity && currentAssociativity == "right") {

			operatorStack = append(operatorStack, string(s))
			continue
		}

		if currentPrecedence < lastPrecedence || (currentPrecedence == lastPrecedence && currentAssociativity == lastAssociativity && currentAssociativity == "left") {
			// TODO O erro ta aqui
			// Isso aqui precisa ter uma verificação pra tirar todo mundo que não bater a precedência corretamente!

			value := operatorStack[len(operatorStack)-1]
			operatorStack[len(operatorStack)-1] = ""
			operatorStack = operatorStack[:len(operatorStack)-1]
			operandStack = append(operandStack, value)
			operatorStack = append(operatorStack, string(s))

			continue
		}

	}

	// Certo, aqui eu preciso adicionar todos os outros que restaram

	operatorStack = slices.DeleteFunc(operatorStack, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})

	if len(operatorStack) >= 1 {
		operandStack = append(operandStack, operatorStack...)
	}

	return operandStack, nil
}

func getPrecedenceInfo(r string) (int, string, error) {
	switch r {
	case "^":
		return 4, "right", nil
	case "*":
		return 3, "left", nil
	case "/":
		return 3, "left", nil
	case "+":
		return 2, "left", nil
	case "-":
		return 2, "left", nil
	default:
		return 0, "", errors.New("invalid operator informed")
	}
}
