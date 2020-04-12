package main

import (
	"io/ioutil"
	"strings"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/parser"

	"github.com/PaulioRandall/scarlet-go/streams/evaluator"
	"github.com/PaulioRandall/scarlet-go/streams/scanner"
)

func main() {

	file := "./test.scarlet"

	errErr := err.Try(func() {
		src, e := ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}

		run(string(src))
	})

	if errErr != nil {
		err.PrintErr(errErr, file)
	}
}

// run executes the input source code.
func run(s string) {

	tokens := analyseScript(s)
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

// analyseScript performs lexical analysis (scans and evaluates) on the script
// s returning an array of tokens.
func analyseScript(s string) []lexeme.Token {
	tokens := scanner.ScanAll(s)
	tokens = evaluator.EvalAll(tokens)
	return tokens
}

// printToken prints a token nicely.
func printToken(tk lexeme.Token) {

	switch k := tk.Lexeme; k {
	case lexeme.LEXEME_TERMINATOR:
		println(k)
	case lexeme.LEXEME_EOF:
		println(k)
		println()
	default:
		print(k + " ")
	}
}

// parseTokens parses the token slice into an executable statement.
func parseTokens(tokens []lexeme.Token) parser.Stat {

	in := make(chan lexeme.Token)

	go func() {
		for _, tk := range tokens {
			in <- tk
		}

		close(in)
	}()

	p := parser.New(in)
	return p.Parse()
}
