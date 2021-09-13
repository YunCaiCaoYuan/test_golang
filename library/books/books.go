package books

import (
	"encoding/json"
	"fmt"
)

type Book struct  {
	Title  string
	Author string
	Pages  int
}

func (b Book) AuthorLastName() (interface{}, interface{}) {
	fmt.Println("AuthorLastName")
	return b.Author,b.Title
}
func (b Book) CategoryByLength() string{
	if b.Pages >300{
		return "NOVEL"
	} else if b.Pages<100 {
		return "SMALL STORY"
	} else {
		return "SHORT STORY"
	}
}
func NewBookFromJSON(json json.Decoder) error {
	book:=json.Decode(json)
	return book
}



