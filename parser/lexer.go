package parser

import (
	"fmt"
	"strconv"
)

// Lex : Lexing given string to tokens
func Lex(str []rune) ([][]rune, error) {
	var tokens [][]rune
	for len(str) > 0 {
		jsonStr, _, err := lexString(str)
		if err != nil {
			return nil, err
		}
		if len(jsonStr) > 0 {
			tokens = append(tokens, jsonStr)
			continue
		}

		char := str[0]
		if contains(char, jsonWhiteSpace) {
			str = str[1:]
		} else if contains(char, jsonSyntax) {
			tokens = append(tokens, []rune{char})
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

func lexString(str []rune) ([]rune, []rune, error) {
	var jsonString []rune

	if str[0] == jsonQuote {
		str = str[1:]
	} else {
		return nil, str, nil
	}

	for _, char := range str {
		if char == jsonQuote {
			return jsonString, str[len(jsonString)+1:], nil
		}
		jsonString = append(jsonString, char)
	}

	return nil, nil, fmt.Errorf("Expected end-of-string quote")
}

func lexBool(str []rune) (res bool, ok bool, rest []rune, err error) {
	lenFalse := len("false")
	lenTrue := len("true")
	if len(str) >= lenFalse && string(str[:lenFalse]) == "false" {
		return false, true, str[lenFalse:], nil
	} else if len(str) >= lenTrue && string(str[:lenTrue]) == "true" {
		return true, true, str[lenTrue:], nil
	} else {
		return false, false, str, nil
	}
}

func lexNumber(str []rune) (interface{}, []rune, error) {
	var jsonNumber []rune
	var numbers = map[rune]struct{}{}
	for i := 0; i < 10; i++ {
		numbers[rune(i+48)] = struct{}{}
	}
	numbers['-'] = struct{}{}
	numbers['e'] = struct{}{}
	numbers['.'] = struct{}{}
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
