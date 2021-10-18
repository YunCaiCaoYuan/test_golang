package main

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

// 注意，读操作r可能会看到并发的写操作w。即使这样也不能表明r之后的读能看到w之前的写。
func main() {
	go f()
	g()
}

