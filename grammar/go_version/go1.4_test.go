package go_version

import (
	"fmt"
	"testing"
)

func Test_range(t *testing.T) {
	list := []int{1, 2, 3, 4}
	for range list {
		fmt.Println("hello")
	}
}
