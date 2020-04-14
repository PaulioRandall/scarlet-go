package main

import (
	"io/ioutil"
	"strings"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/parser"

	"github.com/PaulioRandall/scarlet-go/streams/parser/alpha"
	"github.com/PaulioRandall/scarlet-go/streams/parser/beta"
	"github.com/PaulioRandall/scarlet-go/streams/parser/charlie"
	"github.com/PaulioRandall/scarlet-go/streams/sanitiser"
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

	scannedTokens := scanner.ScanAll(s, scanner.DEFAULT)
	println("***After token scanning***\n")
	lexeme.PrintTokens(scannedTokens)

	sanitisedTokens := sanitiser.SanitiseAll(scannedTokens)
	println("***After token sanitisation***\n")
	lexeme.PrintTokens(sanitisedTokens)

	alphaStats := alpha.TransformAll(sanitisedTokens)
	println("***After alpha statement partitioning***\n")
	alpha.PrintAll(alphaStats)

	betaStats := beta.TransformAll(alphaStats)
	println("***After beta statement partitioning***\n")
	beta.PrintAll(betaStats)

	charlieStats := charlie.TransformAll(betaStats)
	println("***After charlie statement partitioning***\n")
	charlie.PrintAll(charlieStats)

	println("TODO: Parse expression tokens\n")

	exe := parseTokens(sanitisedTokens)

	println(strings.ReplaceAll(exe.String(), "\t", "  "))

	println("\nExecuting...\n")
	ctx := parser.NewContext()
	exe.Eval(ctx)

	println(ctx.String())
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
