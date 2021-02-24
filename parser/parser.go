package parser

import (
	"fmt"
)

func FromString(str string) (interface{}, error) {
	tokens, err := Lex(str)
	if err != nil {
		return nil, err
	}
	result, _, err := parse(tokens)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parse(tokens []interface{}) (interface{}, []interface{}, error) {
	t := tokens[0]

	if t == jsonLeftBrace {
		var object interface{}
		var error error
		object, tokens, error = parseObject(tokens[1:])
		if error != nil {
			return nil, nil, error
		}
		return object, tokens, nil

	} else if t == jsonLeftBracket {
		var array interface{}
		var error error
		array, tokens, error = parseArray(tokens[1:])
		if error != nil {
			return nil, nil, error
		}
		return array, tokens, nil
	}
	return t, tokens[1:], nil
}

func parseArray(tokens []interface{}) ([]interface{}, []interface{}, error) {
	var jsonArray []interface{}
	t := tokens[0]
	if t == jsonRightBracket {
		return jsonArray, tokens[1:], nil
	}

	for {
		var json interface{}
		var err error
		json, tokens, err = parse(tokens)
		if err != nil {
			return nil, nil, err
		}

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
		var err error
		jsonValue, tokens, err = parse(tokens[1:])
		if err != nil {
			return nil, nil, err
		}
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
