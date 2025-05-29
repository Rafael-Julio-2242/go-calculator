package main

import (
	calc "calculator/calc"
	"fmt"
)

func main() {

	input, errInput := calc.Input()

	if errInput != nil {
		fmt.Println("Error in Input! ", errInput)
		return
	}

	result, errResult := calc.Eval(input)

	if errResult != nil {
		fmt.Println("Error on Evaluation expression! ", errResult)
		return
	}

	fmt.Println("Result: ", result)
}
