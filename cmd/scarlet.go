package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/scanner/source"
	"github.com/PaulioRandall/scarlet-go/token"
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

	s := source.New(src)

	t, s, e := s()
	if e != nil {
		panic(e.String())
	}

	for t != token.Empty() {
		print(t.Kind.Name() + " ")
		t, s, e = s()

		if e != nil {
			panic(e.String())
		}
	}
}
