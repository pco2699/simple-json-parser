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
