package main

import (
	"fmt"
	"testing"
)

func Test_int(t *testing.T) {
	lexer := NewSimpleLexer()

	script := "int age = 45;"
	fmt.Println("parse :", script)
	tokenReader := lexer.tokenize(script)

	token := simpleToken{}
	token.Dump(tokenReader)
}

func Test_inta(t *testing.T) {
	lexer := NewSimpleLexer()

	script := "inta age = 45;"
	fmt.Println("\nparse :", script)
	tokenReader := lexer.tokenize(script)

	token := simpleToken{}
	token.Dump(tokenReader)
}

func Test_in(t *testing.T) {
	lexer := NewSimpleLexer()

	script := "in age = 45;"
	fmt.Println("\nparse :", script)
	tokenReader := lexer.tokenize(script)

	token := simpleToken{}
	token.Dump(tokenReader)
}

func Test_GE(t *testing.T) {
	lexer := NewSimpleLexer()

	script := "age >= 45;"
	fmt.Println("\nparse :", script)
	tokenReader := lexer.tokenize(script)

	token := simpleToken{}
	token.Dump(tokenReader)
}

func Test_GT(t *testing.T) {
	lexer := NewSimpleLexer()

	script := "age > 45;"
	fmt.Println("\nparse :", script)
	tokenReader := lexer.tokenize(script)

	token := simpleToken{}
	token.Dump(tokenReader)
}
