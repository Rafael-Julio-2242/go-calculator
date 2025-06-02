package calculator

import (
	"errors"
	"slices"
	"strings"
)

func ShuntingYard(expression string) ([]string, error) {

	var operatorStack []string
	var operandStack []string

	numericRange := "0123456789"
	currentNumber := ""

	for _, s := range expression {
		/*
			fmt.Println()
			fmt.Println("operandStack: ", operandStack)
			fmt.Println("operatorStack: ", operatorStack)
			fmt.Println("[s]: ", string(s))
			fmt.Println("Current Number: ", currentNumber) */

		if strings.Contains(numericRange, string(s)) {
			currentNumber += string(s)
			if rune(expression[len(expression)-1]) == s {
				operandStack = append(operandStack, currentNumber)
				currentNumber = ""
			}
			continue
		} else {
			operandStack = append(operandStack, currentNumber)
			currentNumber = ""
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

			value := operatorStack[len(operatorStack)-1]
			operatorStack[len(operatorStack)-1] = ""
			operatorStack = operatorStack[:len(operatorStack)-1]
			operandStack = append(operandStack, value)

			// Eu acho que aqui, possivelmente, pode dar algum problema alguma hora
			if len(operatorStack) >= 1 {
				for {

					if len(operatorStack) <= 0 {
						break
					}

					currentTop := operatorStack[len(operatorStack)-1]

					currentTopPrecedence, currentTopAssociativity, errCurrentTop := getPrecedenceInfo(operatorStack[len(operatorStack)-1])

					if errCurrentTop != nil {
						return []string{""}, errors.New("error on getting precedence info")
					}

					if currentTopPrecedence > currentPrecedence || (currentTopPrecedence == currentPrecedence && currentTopAssociativity == currentAssociativity && currentAssociativity == "left") {

						operatorStack[len(operatorStack)-1] = ""
						operatorStack = operatorStack[:len(operatorStack)-1]
						operandStack = append(operandStack, currentTop)

					} else {
						break
					}
				}

			}

			operatorStack = append(operatorStack, string(s))

			continue
		}

	}

	operatorStack = slices.DeleteFunc(operatorStack, func(s string) bool {
		return strings.TrimSpace(s) == ""
	})

	if len(operatorStack) >= 1 {

		slices.Reverse(operatorStack)

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
