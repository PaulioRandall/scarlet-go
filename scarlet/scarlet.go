package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	parser "github.com/PaulioRandall/scarlet-go/pkg/parsers/recursive"
	runtime "github.com/PaulioRandall/scarlet-go/pkg/runtime/alpha"
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

	println("# Scanned:")
	tks := scanner.ScanAll(s)
	token.PrettyPrint(tks)

	println("# Parsed:")
	stats := parser.ParseAll(tks)
	statement.Print(stats)

	println("# Executing...")
	ctx := runtime.Run(stats)
	println("...done!\n")

	println(ctx.String())
}
