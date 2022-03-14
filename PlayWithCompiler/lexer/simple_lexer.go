package main

import (
	"bytes"
	"fmt"
)

type SimpleLexer struct {
	tokenText *bytes.Buffer // StringBuffer
	tokens    []*simpleToken
	token     *simpleToken
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
	if this.isAlpha(ch) {
		if ch == 'i' {
			newState = Id_int1
		} else {
			newState = Id
		}
		this.token.setType(Identifier_)
		this.tokenText.WriteByte(ch)
	} else if this.isDigit(ch) {
		newState = IntLiteral__
		this.token.setType(IntLiteral_)
		this.tokenText.WriteByte(ch)
	} else if ch == '>' {
		newState = GT_
		this.token.setType(GT)
		this.tokenText.WriteByte(ch)
	} else if ch == '+' {
		newState = Plus_
		this.token.setType(Plus)
		this.tokenText.WriteByte(ch)
	} else if ch == '-' {
		newState = Minus_
		this.token.setType(Minus)
		this.tokenText.WriteByte(ch)
	} else if ch == '*' {
		newState = Star_
		this.token.setType(Star)
		this.tokenText.WriteByte(ch)
	} else if ch == '/' {
		newState = Slash_
		this.token.setType(Slash)
		this.tokenText.WriteByte(ch)
	} else if ch == ';' {
		newState = SemiColon_
		this.token.setType(SemiColon)
		this.tokenText.WriteByte(ch)
	} else if ch == '(' {
		newState = LeftParen_
		this.token.setType(SemiColon)
		this.tokenText.WriteByte(ch)
	} else if ch == ')' {
		newState = RightParen_
		this.token.setType(RightParen)
		this.tokenText.WriteByte(ch)
	} else if ch == '=' {
		newState = Assignment_
		this.token.setType(Assignment)
		this.tokenText.WriteByte(ch)
	} else {
		newState = Initial
	}

	return newState
}

// 解析字符串，形成Token
func (this *SimpleLexer) tokenize(code string) *simpleTokenReader {
	this.tokens = make([]*simpleToken, 0)

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
		case GT_:
			if ch == '=' {
				this.token.setType(GE)
				state = GE_
			} else {
				state = this.initToken(ch)
			}
		case GE_, Assignment_, Plus_, Minus_, Star_, Slash_, SemiColon_, LeftParen_, RightParen_:
			state = this.initToken(ch)
		case IntLiteral__:
			if this.isDigit(ch) {
				this.tokenText.WriteByte(ch)
			} else {
				state = this.initToken(ch)
			}
		case Id_int1:
			if ch == 'n' {
				state = Id_int2
				this.tokenText.WriteByte(ch)
			} else if this.isDigit(ch) || this.isAlpha(ch) {
				state = Id
				this.tokenText.WriteByte(ch)
			} else {
				state = this.initToken(ch)
			}
		case Id_int2:
			if ch == 't' {
				state = Id_int3
				this.tokenText.WriteByte(ch)
			} else if this.isDigit(ch) || this.isAlpha(ch) {
				state = Id
				this.tokenText.WriteByte(ch)
			} else {
				state = this.initToken(ch)
			}
		case Id_int3:
			if this.isBlank(ch) {
				this.token.setType(Int)
				state = this.initToken(ch)
			} else {
				state = Id
				this.tokenText.WriteByte(ch)
			}
		default:

		}
		if this.tokenText.Len() > 0 {
			this.initToken(ch)
		}
	}

	return &simpleTokenReader{tokens: this.tokens}
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
func (this *simpleToken) setType(typ TokenType) {
	this.typ = typ
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
