package parser

// Parse : Parse tokens
func Parse(tokens []interface{}) (interface{}, interface{}) {
	t := tokens[0]

	if t == jsonLeftBrace {
		return parseObject(tokens)
	} else if t == jsonLeftBracket {
		return parseArray(tokens)
	}
	return nil, tokens
}

func parseArray(tokens []interface{}) (interface{}, []interface{}) {
	return nil, tokens
}

func parseObject(tokens []interface{}) (interface{}, []interface{}) {
	return nil, tokens
}
