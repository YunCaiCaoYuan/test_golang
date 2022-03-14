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
