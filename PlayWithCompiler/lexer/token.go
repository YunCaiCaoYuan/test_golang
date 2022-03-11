package lexer

type Token interface {
	getType() TokenType
	getText() string
	setText(text string)
}
