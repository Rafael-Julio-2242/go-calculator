package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {

	fmt.Println("------ CALCULATOR -------")
	fmt.Println("q - to quit")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Expression: ")
		inputExpression, errInputExp := reader.ReadString('\n')

		if errInputExp != nil {
			fmt.Println("error on input")
			return
		}

		inputExpression = strings.Replace(inputExpression, "\n", "", 1)

		if inputExpression == "q" {
			fmt.Println("quitting...")
			break
		}

		resultExpression, err := ShuntingYard(inputExpression)

		if err != nil {
			fmt.Println("error: ", err)
			return
		}

		expressionTree, err := MountTree(resultExpression)

		if err != nil {
			fmt.Println(err)
			return
		}

		result, evalError := Eval(expressionTree)

		if evalError != nil {
			fmt.Println(evalError)
			return
		}

		fmt.Println("Result: ", result)
		fmt.Println()
	}

}
