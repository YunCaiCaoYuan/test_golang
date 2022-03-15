package main

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

func (c *ASTNodeType) String() string {
	switch *c {
	case Programm:
		return "Programm"
	case IntDeclaration:
		return "IntDeclaration"
	case AssignmentStmt:
		return "AssignmentStmt"
	case Multiplicative:
		return "Multiplicative"
	case Additive:
		return "Additive"
	case Identifier:
		return "Identifier"
	case IntLiteral:
		return "IntLiteral"
	default:
		return ""
	}
}
