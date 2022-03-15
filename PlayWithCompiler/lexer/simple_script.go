package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
)

// 一个简单的脚本解释器
type SimpleScript struct {
	variables map[string]int
	verbose   bool
}

func NewSimpleScript() *SimpleScript {
	return &SimpleScript{
		variables: make(map[string]int),
		verbose:   false,
	}
}

var verbose bool

// 一个简单的REPL
func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p:%v", p)
			debug.PrintStack()
		}
	}()

	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
	fmt.Println("Simple script language!")

	parser := &SimpleParser{}
	script := NewSimpleScript()
	reader := bufio.NewReader(os.Stdin)

	var scriptText string
	fmt.Print("\n>")

	for {
		line0, _ := reader.ReadString('\n')
		line := strings.TrimRight(line0, "\n")
		if line == "exit();" {
			fmt.Println("good bye!")
			break
		}
		scriptText += line + "\n"
		if strings.HasSuffix(line, ";") {
			tree := parser.parse(scriptText)
			if verbose {
				tree.dumpAST(tree, "")
			}

			script.evaluate(tree, "")

			fmt.Print("\n>")

			scriptText = ""
		}
	}
}

// 遍历AST，计算值
func (this *SimpleScript) evaluate(node *SimpleASTNode, indent string) int {
	result := 0
	if verbose {
		getType := node.getType()
		fmt.Println("indent", "Calculating: ", getType.String())
	}

	switch node.getType() {
	case Programm:
		for _, child := range node.getChildren() {
			result = this.evaluate(child, indent)
		}
	case Additive:
		child1 := node.getChildren()[0]
		value1 := this.evaluate(child1, indent+"\t")
		child2 := node.getChildren()[1]
		value2 := this.evaluate(child2, indent+"\t")
		if node.getText() == "+" {
			result = value1 + value2
		} else {
			result = value1 - value2
		}
	case Multiplicative:
		child1 := node.getChildren()[0]
		value1 := this.evaluate(child1, indent+"\t")
		child2 := node.getChildren()[1]
		value2 := this.evaluate(child2, indent+"\t")
		if node.getText() == "*" {
			result = value1 * value2
		} else {
			result = value1 / value2
		}
	case IntLiteral:
		result, _ = strconv.Atoi(node.getText())
	case Identifier:
		varName := node.getText()
		if _, ok := this.variables[varName]; ok {
			value := this.variables[varName]
			result = value
		} else {
			panic(fmt.Errorf("unknown variable: " + varName))
		}
	case AssignmentStmt:
		varName := node.getText()
		if _, ok := this.variables[varName]; !ok {
			panic(fmt.Errorf("unknown variable: " + varName))
		}
	case IntDeclaration:
		varName := node.getText()
		var varValue int
		if len(node.getChildren()) > 0 {
			child := node.getChildren()[0]
			result = this.evaluate(child, indent+"\t")
			varValue = result
		}
		this.variables[varName] = varValue
	default:
	}

	if verbose {
		fmt.Println(indent, "Result: ", result)
	} else if indent == "" {
		if node.getType() == IntDeclaration || node.getType() == AssignmentStmt {
			fmt.Println(node.getText(), ": ", result)
		} else if node.getType() != Programm {
			fmt.Println(result)
		}
	}
	return result
}
