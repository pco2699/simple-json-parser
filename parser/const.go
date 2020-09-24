package parser

const (
	jsonComma        = ','
	jsonColon        = ':'
	jsonLeftBracket  = '['
	jsonRightBracket = ']'
	jsonLeftBrace    = '{'
	jsonRightBrace   = '}'
	jsonQuote        = '"'
)

var jsonWhiteSpace = map[rune]struct{}{
	' ':  {},
	'\t': {},
	'\b': {},
	'\n': {},
	'\r': {},
}

var jsonSyntax = map[rune]struct{}{
	jsonColon:        {},
	jsonComma:        {},
	jsonLeftBrace:    {},
	jsonRightBrace:   {},
	jsonLeftBracket:  {},
	jsonRightBracket: {},
}

var numbers = createNumbers()

func createNumbers() (numbers map[rune]struct{}) {
	numbers = map[rune]struct{}{}
	for i := 0; i < 10; i++ {
		numbers[rune(i+48)] = struct{}{}
	}
	numbers['-'] = struct{}{}
	numbers['e'] = struct{}{}
	numbers['.'] = struct{}{}
	return
}
