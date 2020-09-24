package parser

import (
	"fmt"
	"strconv"
)

// Lex : Lexing given string to tokens
func Lex(str []rune) ([]interface{}, error) {
	var tokens []interface{}
	for len(str) > 0 {
		var jsonStr string
		var err error
		jsonStr, str, err = lexString(str)
		if err != nil {
			return nil, err
		}
		if len(jsonStr) > 0 {
			tokens = append(tokens, jsonStr)
			continue
		}

		var jsonNum interface{}
		jsonNum, str, err = lexNumber(str)
		if err != nil {
			return nil, err
		}
		if jsonNum != nil {
			tokens = append(tokens, jsonNum)
			continue
		}

		var jsonBool bool
		var ok bool
		jsonBool, ok, str = lexBool(str)
		if err != nil {
			return nil, err
		}
		if ok {
			tokens = append(tokens, jsonBool)
			continue
		}

		var jsonNull *interface{}
		jsonNull, ok, str = lexNull(str)
		if err != nil {
			return nil, err
		}
		if ok {
			tokens = append(tokens, jsonNull)
			continue
		}

		char := str[0]
		if contains(char, jsonWhiteSpace) {
			str = str[1:]
		} else if contains(char, jsonSyntax) {
			tokens = append(tokens, char)
			str = str[1:]
		} else {
			return nil, fmt.Errorf("Unexpected character: %s", string(char))
		}
	}
	return tokens, nil
}

func contains(word rune, wordMap map[rune]struct{}) bool {
	_, ok := wordMap[word]
	return ok
}

func lexString(str []rune) (string, []rune, error) {
	var jsonString []rune

	if str[0] == jsonQuote {
		str = str[1:]
	} else {
		return "", str, nil
	}

	for _, char := range str {
		if char == jsonQuote {
			return string(jsonString), str[len(jsonString)+1:], nil
		}
		jsonString = append(jsonString, char)
	}

	return "", nil, fmt.Errorf("Expected end-of-string quote")
}

func lexBool(str []rune) (res bool, ok bool, rest []rune) {
	lenFalse := len("false")
	lenTrue := len("true")
	if len(str) >= lenFalse && string(str[:lenFalse]) == "false" {
		return false, true, str[lenFalse:]
	} else if len(str) >= lenTrue && string(str[:lenTrue]) == "true" {
		return true, true, str[lenTrue:]
	} else {
		return false, false, str
	}
}

func lexNull(str []rune) (res *interface{}, ok bool, rest []rune) {
	lenNull := len("null")
	if len(str) >= lenNull && string(str[:lenNull]) == "null" {
		return nil, true, str[lenNull:]
	}
	return nil, false, str
}

func lexNumber(str []rune) (interface{}, []rune, error) {
	var jsonNumber []rune
	isFloat := false

	for _, char := range str {
		isFloat = char == '.' || isFloat
		if contains(char, numbers) {
			jsonNumber = append(jsonNumber, char)
		} else {
			break
		}
	}

	rest := str[len(jsonNumber):]

	if len(jsonNumber) == 0 {
		return nil, str, nil
	}
	if isFloat {
		f, err := strconv.ParseFloat(string(jsonNumber), 32)
		if err != nil {
			return nil, nil, err
		}
		return float32(f), rest, nil
	}

	i, err := strconv.ParseInt(string(jsonNumber), 10, 32)
	if err != nil {
		return nil, nil, err
	}
	return int32(i), rest, nil
}
