package main

import (
	"io/ioutil"
	"strings"

	"github.com/PaulioRandall/scarlet-go/bard"
	"github.com/PaulioRandall/scarlet-go/lexor"
	"github.com/PaulioRandall/scarlet-go/parser"
	"github.com/PaulioRandall/scarlet-go/token"
)

func main() {

	file := "./test.scarlet"

	src, e := ioutil.ReadFile(file)
	if e != nil {
		panic(e)
	}

	b := bard.NewRhapsodist(file)

	b.CatchNightmare(func() {
		run(string(src))
	})
}

// run executes the input source code.
func run(src string) {

	tokens := collectTokens(src)
	for _, tk := range tokens {
		printToken(tk)
	}

	exe := parseTokens(tokens)

	println(strings.ReplaceAll(exe.String(), "\t", "  "))

	println("\nExecuting...\n")
	ctx := parser.NewContext()
	exe.Eval(ctx)

	println(ctx.String())
}

// collectTokens reads tokens from the 'src' into an array.
func collectTokens(src string) (r []token.Token) {

	var st lexor.TokenStream
	var tk token.Token

	st = lexor.NewScanner(src)
	st = lexor.NewEvaluator(st)

	for tk = st.Next(); tk.Lexeme != token.LEXEME_EOF; tk = st.Next() {
		r = append(r, tk)
	}
	r = append(r, tk)

	return
}

// printToken prints a token nicely.
func printToken(tk token.Token) {

	switch k := tk.Lexeme; k {
	case token.LEXEME_TERMINATOR:
		println(k)
	case token.LEXEME_EOF:
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
