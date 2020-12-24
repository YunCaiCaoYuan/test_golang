package main

import "fmt"

func sqrt(x float64, i float64) (float64, float64) {
	remain := (i*i - x) / (2 * i)
	i = i - remain
	if remain > 0 {
		return sqrt(x, i)
	} else {
		return i, remain
	}
}
func get_sqrt(x float64) float64 {
	i, _ := sqrt(x, x)
	return i
}
func main() {
	fmt.Println(get_sqrt(2))
	fmt.Println(get_sqrt(3))
}
