package calculator

import (
	"errors"
)

func Eval(expression []string) (string, error) {

	var valueA string
	var valueB string
	var signal string

	for i, s := range expression { // Enquanto o resultado não estiver completo, não retorna nada
		// De início, fazer as coisas de forma sequencial
		// Depois a gente adiciona as regras

		if i%2 == 0 { // É um número
			if valueA == "" {
				valueA = s
			} else if valueB == "" {
				valueB = s
			}
		} else { // É um sinal
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
				break
			case "-":
				tempResult, err = Sub(valueA, valueB)

				if err != nil {
					return "", err
				}
				break
			case "*":
				tempResult, err = Multiply(valueA, valueB)

				if err != nil {
					return "", err
				}
				break
			case "/":
				tempResult, err = Divide(valueA, valueB)

				if err != nil {
					return "", err
				}
				break
			default:
				errors.New("Error: Invalid Operation!")
				break
			}

			valueA = tempResult
			valueB = ""
			signal = ""
		}

	}
	return valueA, nil
}
