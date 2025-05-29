package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() ([]string, error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--------- CALCULATOR ---------")

	fmt.Print("Expression: ")
	expression, errInputExpression := reader.ReadString('\n')

	if errInputExpression != nil {
		fmt.Println("Error on Inputing Expression!")
		return []string{}, errInputExpression
	}

	expression = strings.ReplaceAll(expression, "\n", "")

	var valuesExpression []string
	var currentExpressionValue string

	for i, r := range expression {

		if r == ' ' {
			continue
		}

		if r == '+' || r == '-' || r == '*' || r == '/' {
			valuesExpression = append(valuesExpression, currentExpressionValue, string(r))
			currentExpressionValue = ""
		} else if (len(expression) - 1) == i {
			currentExpressionValue += string(r)
			valuesExpression = append(valuesExpression, currentExpressionValue)
			currentExpressionValue = ""
		} else {
			currentExpressionValue += string(r)
		}

	}

	return valuesExpression, nil
}
