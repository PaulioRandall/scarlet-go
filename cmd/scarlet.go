package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/lexor/strimmer"
	"github.com/PaulioRandall/scarlet-go/perror"
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
	var e perror.Perror
	st := strimmer.New(src)

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
		println(k.Name())
	} else {
		print(k.Name() + " ")
	}
}
