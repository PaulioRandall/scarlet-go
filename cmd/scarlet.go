package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/lexor/evaluator"
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

	var t token.Token
	var e lexor.ScanErr
	st := evaluator.New(src)

	for st != nil {
		t, st, e = st()

		if e != nil {
			panic(e.String())
		}

		printToken(t)
	}
}

// Prints the tokens nicely
func printToken(t token.Token) {
	k := t.Kind()

	if k == token.NEWLINE {
		println(k)
	} else {
		print(k + " ")
	}
}
