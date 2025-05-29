package calculator

import (
	"errors"
	"strconv"
)

func VerifyTypes(a string, b string) (string, string, error) {

	var typeA string
	var typeB string

	_, errorConvertIntValue := strconv.ParseInt(a, 10, 64)

	if errorConvertIntValue == nil {
		typeA = "int"
	}

	if typeA != "int" {
		_, errorConvertFloatValue := strconv.ParseFloat(a, 64)

		if errorConvertFloatValue == nil {
			typeA = "float"
		} else {
			err := errors.New("error converting value a")
			return "", "", err
		}
	}

	_, errorConvertIntValue = strconv.Atoi(b)

	if errorConvertIntValue == nil {
		typeB = "int"
	}

	if typeB != "int" {

		_, errorConvertFloatValue := strconv.ParseFloat(b, 64)

		if errorConvertFloatValue == nil {
			typeB = "float"
		} else {
			err := errors.New("error converting value b")
			return "", "", err
		}
	}

	return typeA, typeB, nil
}
