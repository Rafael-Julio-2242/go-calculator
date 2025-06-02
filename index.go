package main

import (
	calc "calculator/calc"
	"fmt"
)

func main() {

	testExpression := "2+2*2/2" // = 10

	resultExpression, err := calc.ShuntingYard(testExpression)

	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	expressionTree, err := calc.MountTree(resultExpression)

	if err != nil {
		fmt.Println(err)
		return
	}

	result, evalError := calc.Eval(expressionTree)

	if evalError != nil {
		fmt.Println(evalError)
		return
	}

	fmt.Println("Result: ", result)

}
