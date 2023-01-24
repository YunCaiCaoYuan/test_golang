package trie

import (
	"fmt"
	"testing"
)

// 字母是边

type Trie struct {
	children []*Trie
	isEnd    bool
	members  []string
}

func Constructor() *Trie {
	return &Trie{
		children: make([]*Trie, 26, 26),
	}
}

func (this *Trie) Insert(word string) {
	node := this

	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = Constructor()
		}
		node = node.children[ch]
	}
	node.isEnd = true

}

func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *Trie) FindStartsWith(prefix string) []string {
	node := this.searchPrefix(prefix)
	this.members = make([]string, 0)
	if node != nil {
		this._findStartsWith(node, prefix)
	}
	return this.members
}

func (this *Trie) _findStartsWith(node *Trie, prefix string) {
	for i := 0; i < 26; i++ {
		if node.children[i] == nil && node.isEnd {
			this.members = append(this.members, prefix)
			return
		}
		if node.children[i] != nil {
			this._findStartsWith(node.children[i], prefix+string(rune(i+int('a'))))
		}
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func Test_trie(t *testing.T) {
	ret := false

	obj := Constructor()
	obj.Insert("apple")
	ret = obj.Search("apple")
	fmt.Println(ret) // true
	ret = obj.Search("app")
	fmt.Println(ret) // false
	ret = obj.StartsWith("app")
	fmt.Println(ret) // true

	obj.Insert("appk")
	res := obj.FindStartsWith("app")
	printArr(res)
}

func printArr(list []string) {
	for _, item := range list {
		fmt.Print(item, " ")
	}
	fmt.Println("")
}
