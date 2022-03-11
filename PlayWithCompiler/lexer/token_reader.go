package lexer

type TokenReader interface {
	Read() Token
	Peek() Token
	Unread()
	GetPosition() int
	SetPosition(int)
}
