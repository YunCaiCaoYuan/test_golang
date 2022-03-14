package main

type ASTNode interface {
	getParent() *ASTNode
	getChildren() []ASTNode
	getType() ASTNodeType
	getText() string
}
