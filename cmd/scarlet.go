package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/parser"
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

	tokens := collectTokens(src)
	for _, tk := range tokens {
		printToken(tk)
	}

	exe := parseTokens(tokens)

	println(exe.String())
}

// collectTokens reads tokens from the 'src' into an array.
func collectTokens(src string) (r []token.Token) {

	var st lexor.TokenStream
	var tk token.Token

	st = lexor.NewScanner(src)
	st = lexor.NewEvaluator(st)

	for tk = st.Next(); tk.Kind != token.EOF; tk = st.Next() {
		r = append(r, tk)
	}
	r = append(r, tk)

	return
}

// printToken prints a token nicely.
func printToken(tk token.Token) {

	switch k := tk.Kind; k {
	case token.TERMINATOR:
		println(k)
	case token.EOF:
		println(k)
		println()
	default:
		print(k + " ")
	}
}

// parseTokens parses the token slice into an executable statement.
func parseTokens(tokens []token.Token) parser.Stat {

	in := make(chan token.Token)

	go func() {
		for _, tk := range tokens {
			in <- tk
		}

		close(in)
	}()

	p := parser.New(in)
	return p.Parse()
}
