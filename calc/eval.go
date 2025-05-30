package calculator

import (
	"errors"
	"fmt"
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

	// Aqui eu preciso ajustar uma forma de buscar pelas operações preferenciais

	specialOps := slices.Contains(expression, "*") || slices.Contains(expression, "/")

	if specialOps {
		var multiply bool
		var divide bool
		var specialOpIndex int

		var negativeA bool
		var negativeB bool

		fmt.Println("----------------------------------------------------- ")

		for i, s := range expression {

			fmt.Println("multiply: ", multiply)
			fmt.Println("divide: ", divide)
			fmt.Println("specialOpIndex: ", specialOpIndex)

			fmt.Println("negativeA: ", negativeA)
			fmt.Println("negativeB: ", negativeB)

			fmt.Println("valueA: ", valueA)
			fmt.Println("valueB: ", valueB)
			fmt.Println("previousExpr: ", previousExpr)
			fmt.Println("s: ", s)

			fmt.Println("expression: ", expression)

			if previousExpr == "-" && (s != "-" && s != "+" && s != "*" && s != "/") && valueA == "" {
				negativeA = true
			}

			if previousExpr == "-" && (s != "-" && s != "+" && s != "*" && s != "/") && valueA != "" && valueB == "" {
				negativeB = true
			}

			if (s == "*" || s == "/") && valueA == "" {

				if previousExpr == "-" || previousExpr == "+" || previousExpr == "*" || previousExpr == "/" {
					return "", errors.New("invalid expression")
				}

				if negativeA {
					valueA = "-" + previousExpr
					negativeA = false
				} else {
					valueA = previousExpr
				}
				valueA = previousExpr
				specialOpIndex = i

				if s == "*" {
					multiply = true
				}

				if s == "/" {
					divide = true
				}

			}

			if (previousExpr == "*" || previousExpr == "/") && valueA != "" {
				if s == "*" || s == "/" || s == "+" {
					return "", errors.New("invalid expression")
				}
				if s == "-" {
					negativeB = true
					continue
				}

				if negativeB {
					valueB = "-" + s
				} else {
					valueB = s
				}
			}

			if valueA != "" && valueB != "" {

				if multiply {
					result, errMult := Multiply(valueA, valueB)

					if errMult != nil {
						return "", errors.New("error on multiplication")
					}

					// Pos do valueA é specialOpIndex - 1
					// Preciso remover as coisas dentro das posições specialOpIndex e specialOpIndex + 1

					expression = slices.Delete(expression, specialOpIndex, specialOpIndex+2)
					fmt.Println("MULTIPLIED EXPRESSION 1: ", expression)

					expression = slices.Replace(expression, specialOpIndex-1, specialOpIndex, result)

					fmt.Println("MULTIPLIED EXPRESSION 2: ", expression)

					multiply = false
					valueA = ""
					valueB = ""
				}

				if divide {
					result, errMult := Divide(valueA, valueB)

					if errMult != nil {
						return "", errors.New("error on multiplication")
					}

					// Pos do valueA é specialOpIndex - 1
					// Preciso remover as coisas dentro das posições specialOpIndex e specialOpIndex + 1

					expression = slices.Delete(expression, specialOpIndex, specialOpIndex+1)
					expression = slices.Replace(expression, specialOpIndex-1, specialOpIndex-1, result)

					divide = false
					valueA = ""
					valueB = ""
				}

			}

			previousExpr = s
			fmt.Println("------------------------------------------------- ")

		}
		valueA = ""
		valueB = ""
		previousExpr = ""

	}

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
