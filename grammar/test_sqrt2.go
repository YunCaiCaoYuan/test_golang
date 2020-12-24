package main

import "fmt"
import "unsafe"

func get_sqrt(x float32) float32 {
	xhalf := 0.5 * x
	var i int32 = *(*int32)(unsafe.Pointer(&x))
	i = 0x5f375a86 - (i >> 1)
	x = *(*float32)(unsafe.Pointer(&i))

	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)

	return 1 / x
}

func main() {
	fmt.Println(get_sqrt(4))
}
