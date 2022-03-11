package lexer

type ASTNodeType int

const (
	Programm ASTNodeType = iota + 1

	IntDeclaration
	ExpressionStmt
	AssignmentStmt

	Primary
	Multiplicative
	Additive

	Identifier
	IntLiteral
)
