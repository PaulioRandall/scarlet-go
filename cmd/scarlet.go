package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/lexor/strimmer"
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

	st := strimmer.New(src)

	t, st, e := st()
	if e != nil {
		panic(e.String())
	}

	for t != token.Empty() {
		print(t.Kind().Name() + " ")
		t, st, e = st()

		if e != nil {
			panic(e.String())
		}
	}
}
