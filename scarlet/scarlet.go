package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	statement "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	token "github.com/PaulioRandall/scarlet-go/pkg/z_token"

	parser "github.com/PaulioRandall/scarlet-go/pkg/parsers/z_recursive"
	runtime "github.com/PaulioRandall/scarlet-go/pkg/runtime/z_alpha"
	sanitiser "github.com/PaulioRandall/scarlet-go/pkg/sanitisers/standard"
	scanner "github.com/PaulioRandall/scarlet-go/pkg/scanners/z_matching"
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

	println("# Scanning:")
	tks := scanner.ScanAll(s)
	token.PrettyPrint(tks)

	println("# Sanitising:")
	tks = sanitiser.SanitiseAll(tks)
	token.PrettyPrint(tks)

	println("# Parsing:")
	stats := parser.ParseAll(tks)
	statement.Print(stats)

	println("# Executing...")
	ctx := runtime.Run(stats)
	println("...done!\n")

	println(ctx.String())
}
