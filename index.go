package main

import (
	calc "calculator/calc"
	"fmt"
)

func main() {

	testExpression := "2+2*2/2"

	resultExpression, err := calc.ShuntingYard(testExpression)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Expected expression: 2 2 2 * 2 / +")
	fmt.Println("Result Expression: ", resultExpression)
}
