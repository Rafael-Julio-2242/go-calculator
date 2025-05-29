package calculator

import (
	"fmt"
	"strconv"
)

func Divide(a string, b string) (string, error) {

	valueA, errConvertingA := strconv.ParseFloat(a, 64)
	valueB, errConvertingB := strconv.ParseFloat(b, 64)

	if errConvertingA != nil {
		return "", errConvertingA
	}

	if errConvertingB != nil {
		return "", errConvertingB
	}

	result := valueA / valueB

	formatedResult := fmt.Sprintf("%.2f", result)

	return formatedResult, nil
}
