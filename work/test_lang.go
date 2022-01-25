package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"strings"
)

// Language .
type Language struct {
	Tag language.Tag
	p   *message.Printer
	//cache *ristretto.Cache
	Code string
}

func main() {
	tag := language.TraditionalChinese
	parts := strings.SplitN(tag.String(), "-u-rg", 2)
	fmt.Println("parts=", parts)

	Languages := map[string]*Language{}
	l := &Language{
		Tag:  tag,
		Code: parts[0],
		p:    message.NewPrinter(tag),
		//cache: cache,
	}
	Languages[parts[0]] = l

	for item := range Languages {
		fmt.Println("item=", item)
	}

	ll := make([]int, 0)
	ll = append(ll, 123, 456, 789)
	for item := range ll {
		fmt.Println("item123=", item)
	}
}
