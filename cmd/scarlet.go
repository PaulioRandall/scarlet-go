package main

import (
	"io/ioutil"
	//"strings"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/parser"

	"github.com/PaulioRandall/scarlet-go/streams/parser/recursive"
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

	statements := recursive.ParseAll(sanitisedTokens)
	println("***After alpha statement partitioning***\n")
	recursive.Print(statements)

	exe := parseTokens(sanitisedTokens)

	//println(strings.ReplaceAll(exe.String(), "\t", "  "))

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
