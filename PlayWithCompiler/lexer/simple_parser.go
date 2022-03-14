package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// 语法解析器
type SimpleParser struct{}

func (this *SimpleParser) parse(script string) *SimpleASTNode {
	lexer := new(SimpleLexer)
	tokens := lexer.tokenize(script)
	rootNode := this.prog(tokens)
	return rootNode
}

func (this *SimpleParser) prog(tokens *simpleTokenReader) *SimpleASTNode {
	node := NewSimpleASTNode(Programm, "pwc")
	for tokens.Peek() != nil {
		child := this.intDeclare(tokens)
		if child == nil {
			child = this.expressionStatement(tokens)
		}
		if child == nil {
			child = this.assignmentStatement(tokens)
		}
		if child != nil {
			node.addChild(child)
		} else {
			panic("unknown statement")
		}
	}
	return node
}

// 表达式语句
func (this *SimpleParser) expressionStatement(tokens *simpleTokenReader) *SimpleASTNode {
	pos := tokens.GetPosition()
	node := this.additive(tokens)
	if node != nil {
		token := tokens.Peek()
		if token != nil && token.getType() == SemiColon {
			tokens.Read()
		} else {
			node = nil
			tokens.SetPosition(pos)
		}
	}
	return node
}

// 赋值语句
func (this *SimpleParser) assignmentStatement(tokens *simpleTokenReader) *SimpleASTNode {
	var node *SimpleASTNode
	token := tokens.Peek()
	if token != nil && token.getType() == Identifier_ {
		token := tokens.Read()
		node = NewSimpleASTNode(AssignmentStmt, token.getText())
		token = tokens.Peek()
		if token != nil && token.getType() == Assignment {
			tokens.Read()
			child := this.additive(tokens)
			if child == nil {
				panic("invalide assignment statement, expecting an expression")
			} else {
				node.addChild(child)
				token = tokens.Peek()
				if token != nil && token.getType() == SemiColon {
					tokens.Read()
				} else {
					panic("invalid statement, expecting semicolon")
				}
			}
		} else {
			tokens.Unread()
			node = nil
		}
	}
	return node
}

func (this *SimpleParser) intDeclare(tokens *simpleTokenReader) *SimpleASTNode {
	var node *SimpleASTNode
	token := tokens.Peek()
	if token != nil && token.getType() == Int {
		token = tokens.Read()
		if tokens.Peek().getType() == Identifier_ {
			token = tokens.Read()
			node := NewSimpleASTNode(IntDeclaration, token.getText())
			token = tokens.Peek()
			if token != nil && token.getType() == Assignment {
				tokens.Read()
				child := this.additive(tokens)
				if child == nil {
					panic("invalide variable initialization, expecting an expression")
				} else {
					node.addChild(child)
				}
			}
		} else {
			panic("variable name expected")
		}

		if node != nil {
			token = tokens.Peek()
			if token != nil && token.getType() == SemiColon {
				tokens.Read()
			} else {
				panic("invalid statement, expecting semicolon")
			}
		}
	}
	return node
}

// 加法表达式
func (this *SimpleParser) additive(tokens *simpleTokenReader) *SimpleASTNode {
	child1 := this.multiplicative(tokens)
	node := child1
	if child1 != nil {
		for {
			token := tokens.Peek()
			if token != nil && (token.getType() == Plus || token.getType() == Minus) {
				token = tokens.Read()
				child2 := this.multiplicative(tokens)
				if child2 != nil {
					node = NewSimpleASTNode(Additive, token.getText())
					node.addChild(child1)
					node.addChild(child2)
					child1 = node
				} else {
					panic("invalid additive expression, expecting the right part.")
				}
			} else {
				break
			}
		}
	}
	return node
}

// 乘法表达式
func (this *SimpleParser) multiplicative(tokens *simpleTokenReader) *SimpleASTNode {
	child1 := this.primary(tokens)
	node := child1
	for {
		token := tokens.Peek()
		if token != nil && (token.getType() == Star || token.getType() == Slash) {
			token = tokens.Read()
			child2 := this.primary(tokens)
			if child2 != nil {
				node := NewSimpleASTNode(Multiplicative, token.getText())
				node.addChild(child1)
				node.addChild(child2)
				child1 = node
			} else {
				panic("invalid multiplicative expression, expecting the right part.")
			}
		} else {
			break
		}
	}

	return node
}

// 基础表达式
func (this *SimpleParser) primary(tokens *simpleTokenReader) *SimpleASTNode {
	var node *SimpleASTNode
	token := tokens.Peek()
	if token != nil {
		if token.getType() == IntLiteral_ {
			token = tokens.Read()
			node = NewSimpleASTNode(IntLiteral, token.getText())
		} else if token.getType() == Identifier_ {
			token = tokens.Read()
			node = NewSimpleASTNode(Identifier, token.getText())
		} else if token.getType() == LeftParen {
			tokens.Read()
			node = this.additive(tokens)
			if node != nil {
				token = tokens.Peek()
				if token != nil && token.getType() == RightParen {
					tokens.Read()
				} else {
					panic("expecting right parenthesis")
				}
			} else {
				panic("expecting an additive expression inside parenthesis")
			}
		}
	}

	return node
}

// 一个简单的AST节点
type SimpleASTNode struct {
	parent        *SimpleASTNode
	children      []*SimpleASTNode
	cloneChildren []*SimpleASTNode
	nodeType      ASTNodeType
	text          string
}

func NewSimpleASTNode(nodeType ASTNodeType, text string) *SimpleASTNode {
	return &SimpleASTNode{
		nodeType: nodeType,
		text:     text,
	}
}
func (this *SimpleASTNode) getParent() *SimpleASTNode {
	return this.parent
}
func (this *SimpleASTNode) getChildren() []*SimpleASTNode {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(&this.children); err != nil {
		panic("getChildren" + err.Error())
	}
	if err := gob.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(&this.cloneChildren); err != nil {
		panic("getChildren" + err.Error())
	}
	return this.cloneChildren
}
func (this *SimpleASTNode) getType() ASTNodeType {
	return this.nodeType
}
func (this *SimpleASTNode) getText() string {
	return this.text
}
func (this *SimpleASTNode) addChild(child *SimpleASTNode) {
	this.children = append(this.children, child)
	child.parent = this
	return
}

func (this *SimpleASTNode) dumpAST(node *SimpleASTNode, indent string) {
	fmt.Println("indent", node.getType(), " ", node.getText())
	for _, child := range node.getChildren() {
		this.dumpAST(child, indent+"\t")
	}
}
