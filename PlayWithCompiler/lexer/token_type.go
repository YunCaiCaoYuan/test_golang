package main

type TokenType int

const (
	Plus TokenType = iota + 1
	Minus
	Star
	Slash

	GE
	GT
	EQ
	LE
	LT

	SemiColon
	LeftParen
	RightParen

	Assignment

	If
	Else

	Int

	Identifier_

	IntLiteral_

	StringLiteral
)

func (c *TokenType) String() string {
	switch *c {
	case Plus:
		return "Plus"
	case Minus:
		return "Minus"
	case Star:
		return "Star"
	case Slash:
		return "Slash"

	case GE:
		return "GE"
	case GT:
		return "GT"

	case SemiColon:
		return "SemiColon"
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case Assignment:
		return "Assignment"

	case Int:
		return "Int"
	case Identifier_:
		return "Identifier_"
	case IntLiteral_:
		return "IntLiteral_"

	default:
		return "unkown"
	}
}
