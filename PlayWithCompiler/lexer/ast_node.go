package lexer

type ASTNode interface {
	getParent() *ASTNode
	getChildren() []ASTNode
	getType() ASTNodeType
	getText() string
}
