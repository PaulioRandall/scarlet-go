package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/scanner"
)

func main() {

	b, e := ioutil.ReadFile("./test.scarlet")
	if e != nil {
		panic(e)
	}

	s := scanner.New(string(b))
	tok, s, pe := s()
	if pe != nil {
		panic(pe.String())
	}

	for tok != scanner.EmptyTok() {
		print(tok.Kind.Name() + " ")
		tok, s, pe = s()

		if pe != nil {
			panic(pe.String())
		}
	}
}
