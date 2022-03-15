package main

import (
	"fmt"
	"strconv"
)

type SimpleCalculator struct{}

func (this *SimpleCalculator) evaluate(script string) {
	tree := this.parse(script)
	this.evaluate2(tree, "")
}

func (this *SimpleCalculator) parse(code string) *SimpleASTNode {
	lexer := &SimpleLexer{}
	tokens := lexer.tokenize(code)
	rootNode := this.prog(tokens)
	return rootNode
}

func (this *SimpleCalculator) evaluate2(node *SimpleASTNode, indent string) int {
	result := 0
	fmt.Println(indent, "Calculating: ", node.getType())
	switch node.getType() {
	case Programm:
		for _, child := range node.getChildren() {
			result = this.evaluate2(child, indent+"\t")
		}
	case Additive:
		child1 := node.getChildren()[0]
		value1 := this.evaluate2(child1, indent+"\t")
		child2 := node.getChildren()[1]
		value2 := this.evaluate2(child2, indent+"\t")
		if node.getText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case Multiplicative:
		child1 := node.getChildren()[0]
		value1 := this.evaluate2(child1, indent+"\t")
		child2 := node.getChildren()[1]
		value2 := this.evaluate2(child2, indent+"\t")
		if node.getText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case IntLiteral:
		result, _ = strconv.Atoi(node.getText())
	default:

	}

	fmt.Println(indent, "Result: ", result)
	return result
}

func (this *SimpleCalculator) prog(tokens *simpleTokenReader) *SimpleASTNode {
	node := NewSimpleASTNode(Programm, "Calculator")
	child := this.additive(tokens)
	if child != nil {
		node.addChild(child)
	}
	return node
}

func (this *SimpleCalculator) intDeclare(tokens *simpleTokenReader) *SimpleASTNode {
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

func (this *SimpleCalculator) additive(tokens *simpleTokenReader) *SimpleASTNode {
	child1 := this.multiplicative(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.getType() == Plus || token.getType() == Minus {
			token = tokens.Read()
			child2 := this.additive(tokens)
			if child2 != nil {
				node = NewSimpleASTNode(Additive, token.getText())
				node.addChild(child1)
				node.addChild(child2)
			}
		}
	}
	return node
}

func (this *SimpleCalculator) multiplicative(tokens *simpleTokenReader) *SimpleASTNode {
	child1 := this.primary(tokens)
	node := child1
	token := tokens.Peek()
	if child1 != nil && token != nil {
		if token.getType() == Star || token.getType() == Slash {
			token = tokens.Read()
			child2 := this.multiplicative(tokens)
			if child2 != nil {
				node := NewSimpleASTNode(Multiplicative, token.getText())
				node.addChild(child1)
				node.addChild(child2)
			} else {
				panic("invalid multiplicative expression, expecting the right part.")
			}
		}
	}
	return node
}

func (this *SimpleCalculator) primary(tokens *simpleTokenReader) *SimpleASTNode {
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
