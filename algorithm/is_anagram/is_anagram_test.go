package is_anagram

import (
	"fmt"
	"reflect"
	"testing"
)

func isAnagram(s string, t string) bool {
	sLetterMap := make(map[byte]int)
	for _, item := range []byte(s) {
		sLetterMap[item] += 1
	}
	tLetterMap := make(map[byte]int)
	for _, item := range []byte(t) {
		tLetterMap[item] += 1
	}
	return reflect.DeepEqual(sLetterMap, tLetterMap)
}

func Test_isAnagram(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func Test_isAnagram2(t *testing.T) {
	fmt.Println(isAnagram("rat", "car"))
}
