package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_trim(t *testing.T) {
	bindStr := "+8618780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}

func Test_trim2(t *testing.T) {
	bindStr := "+18780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}

func Test_trim3(t *testing.T) {
	bindStr := "18780170596"
	bindReal := strings.Trim(bindStr, "+")
	fmt.Println(bindReal)
}
