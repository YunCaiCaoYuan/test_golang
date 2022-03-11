package lexer

import (
	"bytes"
	"fmt"
)

type SimpleLexer struct {
	tokenText *bytes.Buffer // StringBuffer
	tokens    []Token
	token     Token
}

func (this *SimpleLexer) isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func (this *SimpleLexer) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (this *SimpleLexer) isBlank(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func (this *SimpleLexer) New() {
	this.token = &simpleToken{}
}

// 有限状态机进入初始状态
func (this *SimpleLexer) initToken(ch byte) DfaState {
	if this.tokenText.Len() > 0 {
		this.token.setText(this.tokenText.String())
		this.tokens = append(this.tokens, this.token)

		this.tokenText = new(bytes.Buffer)
		this.token = &simpleToken{}
	}

	newState := Initial
	// todo

	return newState
}

// 解析字符串，形成Token
func (this *SimpleLexer) tokenize(code string) *simpleTokenReader {
	this.tokens = make([]Token, 0)

	reader := bytes.NewBufferString(code)
	state := Initial
	var ch byte
	var err error
	for ch, err = reader.ReadByte(); err == nil; ch, err = reader.ReadByte() {
		switch state {
		case Initial:
			state = this.initToken(ch)
		case Id:
			if this.isAlpha(ch) || this.isDigit(ch) {
				this.tokenText.WriteByte(ch)
			} else {
				state = this.initToken(ch)
			}
		}
		// todo

	}

	return nil
}

// Token的一个简单实现
type simpleToken struct {
	typ  TokenType
	text string
}

func (this *simpleToken) getType() TokenType {
	return this.typ
}
func (this *simpleToken) getText() string {
	return this.text
}
func (this *simpleToken) setText(text string) {
	this.text = text
}

func (this *simpleToken) Dump(tokenReader simpleTokenReader) {
	fmt.Println("text\ttype")
	var token Token
	for token = tokenReader.Read(); token == nil; token = tokenReader.Read() {
		fmt.Println(token.getText(), "\t\t", token.getType())
	}
}

type DfaState int

const (
	Initial DfaState = iota + 1

	If_
	Id_if1
	Id_if2
	Else_
	Id_else1
	Id_else2
	Id_else3
	Id_else4
	Int_
	Id_int1
	Id_int2
	Id_int3
	Id
	GT_
	GE_

	Assignment_

	Plus_
	Minus_
	Star_
	Slash_

	SemiColon_
	LeftParen_
	RightParen_

	IntLiteral__
)

// 一个简单的Token流
type simpleTokenReader struct {
	tokens []*simpleToken // List<Token>
	pos    int
}

func (this *simpleTokenReader) SimpleTokenReader() {
	this.tokens = make([]*simpleToken, 0)
}
func (this *simpleTokenReader) Read() *simpleToken {
	this.pos++
	if this.pos < len(this.tokens) {
		return this.tokens[this.pos]
	}
	return nil
}
func (this *simpleTokenReader) Peek() *simpleToken {
	if this.pos < len(this.tokens) {
		return this.tokens[this.pos]
	}
	return nil
}
func (this *simpleTokenReader) Unread() {
	if this.pos > 0 {
		this.pos--
	}
}
func (this *simpleTokenReader) GetPosition() int {
	return this.pos
}
func (this *simpleTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(this.tokens) {
		this.pos = pos
	}
}
