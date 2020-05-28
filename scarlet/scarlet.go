package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	parser "github.com/PaulioRandall/scarlet-go/pkg/parsers/recursive"
	runtime "github.com/PaulioRandall/scarlet-go/pkg/runtime/alpha"
	sanitiser "github.com/PaulioRandall/scarlet-go/pkg/sanitisers/standard"
	scanner "github.com/PaulioRandall/scarlet-go/pkg/scanners/matching"
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
		err.Print(os.Stdout, e, file)
	}
}

func run(s string) {

	println("# Scanning...")
	tks := scanner.ScanAll(s)
	token.PrettyPrint(tks)

	println()
	println("# Sanitising...")
	tks = sanitiser.SanitiseAll(tks)
	token.PrettyPrint(tks)

	println()
	println("# Parsing...")
	stats := parser.ParseAll(tks)
	statement.Print(stats)

	println()
	println("# Executing...")
	ctx := runtime.Run(stats)

	println()
	println(ctx.String())
}
