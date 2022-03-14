package lexer

type Token interface {
	getType() TokenType
	getText() string
	setType(TokenType)
	setText(text string)
}
