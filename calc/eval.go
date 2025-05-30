package calculator

import (
	"errors"
	"slices"
	"strings"
)

func Eval(expression []string) (string, error) {

	var valueA string
	var valueB string
	var signal string

	var negative bool
	var previousExpr string

	expression = slices.DeleteFunc(expression, func(exp string) bool {
		return strings.TrimSpace(exp) == ""
	})

	for _, s := range expression { // Enquanto o resultado não estiver completo, não retorna nada
		// De início, fazer as coisas de forma sequencial
		// Depois a gente adiciona as regras

		if s != "+" && s != "-" && s != "*" && s != "/" {
			if valueA == "" {
				if negative || previousExpr == "-" {
					valueA = "-" + s
					negative = false
				} else {
					valueA = s
				}
			} else if valueB == "" {

				if negative || previousExpr == "-" {
					valueB = "-" + s
					negative = false
				} else {
					valueB = s
				}
			}
		} else if previousExpr == "-" && s == "-" {
			signal = "+"
		} else if (previousExpr == "*" || previousExpr == "/" || previousExpr == "+") && s == "-" {
			signal = previousExpr
			negative = true
		} else {
			signal = s
		}

		if valueA != "" && valueB != "" && signal != "" {
			var tempResult string
			var err error

			switch signal {
			case "+":
				tempResult, err = Sum(valueA, valueB)

				if err != nil {
					return "", err
				}
			case "-":
				tempResult, err = Sub(valueA, valueB)

				if err != nil {
					return "", err
				}
			case "*":
				tempResult, err = Multiply(valueA, valueB)

				if err != nil {
					return "", err
				}
			case "/":
				tempResult, err = Divide(valueA, valueB)

				if err != nil {
					return "", err
				}
			default:
				err = errors.New("error: invalid operation")
				return "", err
			}

			valueA = tempResult
			valueB = ""
			signal = ""
		}

		previousExpr = s
	}
	return valueA, nil
}
