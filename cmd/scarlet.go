package main

import (
	"io/ioutil"
	"strings"

	"github.com/PaulioRandall/scarlet-go/err"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/parser"

	"github.com/PaulioRandall/scarlet-go/streams/articulator"
	"github.com/PaulioRandall/scarlet-go/streams/evaluator"
	"github.com/PaulioRandall/scarlet-go/streams/partitioner"
	"github.com/PaulioRandall/scarlet-go/streams/scanner"
	"github.com/PaulioRandall/scarlet-go/streams/token"
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

	scannedTokens := scanner.ScanAll(s)
	println("***After token scanning***\n")
	token.PrintAll(scannedTokens)

	evaluatedTokens := evaluator.EvalAll(scannedTokens)
	println("***After token evaluation***\n")
	token.PrintAll(evaluatedTokens)

	stats := partitioner.PartitionAll(evaluatedTokens)
	println("***After statement partitioning***\n")
	partitioner.PrintAll(stats)

	arts := articulator.ArticulateAll(stats)
	println("***After articulation***\n")
	articulator.PrintAll(arts)

	println("THEN: Parse assignment tokens\n")
	println("THEN: Parse expression tokens\n")

	exe := parseTokens(evaluatedTokens)

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
