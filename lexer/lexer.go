package lexer

func Lex(str string) ([]string, error) {
	var tokens []string
	for len(str) > 0 {
		json_str, str := lex_string(str)
		if len(json_str) > 0 {
			tokens = append(tokens, json_str)
			continue
		}
	}
}

func lex_string(str string) (string, string) {
	return "", str
}
