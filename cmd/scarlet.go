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

	run(string(b))
}

// run executes the input source code.
func run(src string) {

	s := scanner.New(src)

	tok, s, e := s()
	if e != nil {
		panic(e.String())
	}

	for tok != scanner.EmptyTok() {
		print(tok.Kind.Name() + " ")
		tok, s, e = s()

		if e != nil {
			panic(e.String())
		}
	}
}
