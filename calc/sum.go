package calculator

import (
	"fmt"
	"strconv"
)

func Sum(a string, b string) (string, error) {

	typeA, typeB, typesErr := VerifyTypes(a, b)

	if typesErr != nil {
		return "", typesErr
	}

	if typeA == typeB && typeA == "int" {

		intA, _ := strconv.Atoi(a)
		intB, _ := strconv.Atoi(b)

		result := intA + intB

		formatedResult := fmt.Sprintf("%d", result)

		return formatedResult, nil
	}

	valueA, _ := strconv.ParseFloat(a, 64)
	valueB, _ := strconv.ParseFloat(b, 64)

	result := valueA + valueB

	formatedValue := fmt.Sprintf("%.2f", result)

	return formatedValue, nil
}
