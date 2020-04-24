package main

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers"
	"github.com/PaulioRandall/scarlet-go/pkg/runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/sanitiser"
	"github.com/PaulioRandall/scarlet-go/pkg/scanners"
)

func main() { // Run it with `./godo run`

	file := "./test.scarlet"

	e := err.Try(func() {

		src, e := ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}

		run(string(src))
	})

	if e != nil {
		err.PrintErr(e, file)
	}
}

// run executes the input source code.
func run(s string) {

	println("# Scanned:")
	scannedTokens := scanners.ScanAll(s, scanners.DEFAULT)
	token.PrintTokens(scannedTokens)

	println("# Sanitised:")
	sanitisedTokens := sanitiser.SanitiseAll(scannedTokens)
	token.PrintTokens(sanitisedTokens)

	println("# Parsed:")
	statements := parsers.ParseAll(sanitisedTokens, parsers.DEFAULT)
	statement.Print(statements)

	print("# Executing...")
	ctx := runtime.Run(statements)
	println("...done!\n")

	println(ctx.String())
}
