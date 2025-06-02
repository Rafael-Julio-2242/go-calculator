package main

import (
	calc "calculator/calc"
	"fmt"
)

func main() {

	testExpression := "8+2*5/2-3"
	expectedExpression := "8 2 5 * 2 / + 3 -"

	resultExpression, err := calc.ShuntingYard(testExpression)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Expected expression: ", expectedExpression)
	fmt.Println("Result Expression: ", resultExpression)
}
