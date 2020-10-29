package token

import (
	"fmt"
)

type LexList struct {
	Items []Lexeme
	Size  int // Length of the item slice
}

func NewLexList(size, capa int) *LexList {
	return &LexList{
		Items: make([]Lexeme, size, capa),
		Size:  size,
	}
}

func LexListFrom(lexs ...Lexeme) *LexList {
	size := len(lexs)
	items := make([]Lexeme, size, size)
	for i, v := range lexs {
		items[i] = v
	}
	return &LexList{
		Items: items,
		Size:  size,
	}
}

func (ll *LexList) Length() int {
	return ll.Size
}

func (ll *LexList) Capacity() int {
	return cap(ll.Items)
}

func (ll *LexList) Append(l Lexeme) {
	ll.Items = append(ll.Items, l)
	ll.Size++
}

func (ll *LexList) Get(i int) Lexeme {
	ll.checkRange(i)
	return ll.Items[i]
}

func (ll *LexList) Set(i int, l Lexeme) {
	ll.checkRange(i)
	ll.Items[i] = l
}

func (ll *LexList) checkRange(i int) {
	if i < 0 {
		e := fmt.Errorf("Out of range: index '%d' is negative", i)
		panic(e)
	}
	if i >= ll.Size {
		e := fmt.Errorf(
			"Out of range: index '%d' matches or exceeds size %d", i, ll.Size)
		panic(e)
	}
}
