package main

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

// 即使main看到了g != nil并且退出了循环，也不能保证看到g.msg的初始化值。
func main() {
	go setup()
	for g == nil {
	}
	print(g.msg)
}
