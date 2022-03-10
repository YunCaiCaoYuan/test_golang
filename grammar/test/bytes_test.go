package test

import (
	"bytes"
	"fmt"
	"testing"
)

// for循环遍历bytes
func Test_bytes_loop(t *testing.T) {
	code := "lalala"
	reader := bytes.NewBufferString(code)
	var ch byte
	var err error
	for ch, err = reader.ReadByte(); err == nil; ch, err = reader.ReadByte() {
		fmt.Printf("%c\n", ch)
	}
}
