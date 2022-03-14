package main

type Token interface {
	getType() TokenType
	getText() string
	setType(TokenType)
	setText(text string)
}
