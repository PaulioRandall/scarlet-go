package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/lexor"
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

	var st lexor.TokenStream

	st = lexor.NewScanner(src)
	st = lexor.NewEvaluator(st)

	for t := st.Next(); t != (token.Token{}); t = st.Next() {
		if st != nil && t != (token.Token{}) {
			printToken(t)
		}
	}
}

// Prints the tokens nicely
func printToken(t token.Token) {
	k := t.Kind

	if k == token.NEWLINE {
		println(k)
	} else {
		print(k + " ")
	}
}
