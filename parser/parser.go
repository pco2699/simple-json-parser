package parser

import (
	"fmt"
	"os"
)

// Parse : Parse tokens
func Parse(tokens []interface{}) (interface{}, []interface{}) {
	t := tokens[0]

	if t == jsonLeftBrace {
		var object interface{}
		var error error
		object, tokens, error = parseObject(tokens[1:])
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
		return object, tokens

	} else if t == jsonLeftBracket {
		var array interface{}
		var error error
		array, tokens, error = parseArray(tokens[1:])
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
		return array, tokens
	}
	return t, tokens[1:]
}

func parseArray(tokens []interface{}) ([]interface{}, []interface{}, error) {
	var jsonArray []interface{}
	t := tokens[0]
	if t == jsonRightBracket {
		return jsonArray, tokens[1:], nil
	}

	for {
		var json interface{}
		json, tokens = Parse(tokens)

		jsonArray = append(jsonArray, json)

		if t, ok := tokens[0].(rune); ok {
			if t == jsonRightBracket {
				return jsonArray, tokens[1:], nil
			} else if t != jsonComma {
				return nil, nil, fmt.Errorf("Expected comma after object in array")
			} else {
				tokens = tokens[1:]
			}
		}
	}
}

func parseObject(tokens []interface{}) (map[string]interface{}, []interface{}, error) {
	jsonObject := make(map[string]interface{})

	t := tokens[0]
	if t == jsonRightBrace {
		return jsonObject, tokens[1:], nil
	}

	for {
		jsonKey := tokens[0]
		if _, ok := jsonKey.(string); ok {
			tokens = tokens[1:]
		} else {
			return nil, nil, fmt.Errorf("Expected string key, got: %v", jsonKey)
		}
		jsonKeyStr := jsonKey.(string)

		if tokens[0] != jsonColon {
			return nil, nil, fmt.Errorf("Expected colon after key in object, got: %v", t)
		}

		var jsonValue interface{}
		jsonValue, tokens = Parse(tokens[1:])
		jsonObject[jsonKeyStr] = jsonValue

		t = tokens[0]
		if t == jsonRightBrace {
			return jsonObject, tokens[1:], nil
		} else if t != jsonComma {
			return nil, nil, fmt.Errorf("Expected comma after pair in object, got: %v", t)
		}

		tokens = tokens[1:]
	}
}
