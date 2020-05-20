package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
	"github.com/PaulioRandall/scarlet-go/pkg/statement"
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	"github.com/PaulioRandall/scarlet-go/pkg/parsers"
	"github.com/PaulioRandall/scarlet-go/pkg/runtime"
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
		err.Print(os.Stdout, e, file)
	}
}

func run(s string) {

	println("# Scanned:")
	tokens := scanners.ScanAll(s, scanners.DEFAULT)
	token.PrettyPrint(tokens)

	println("# Parsed:")
	statements := parsers.ParseAll(tokens, parsers.DEFAULT)
	statement.Print(statements)

	println("# Executing...")
	ctx := runtime.Run(statements, runtime.DEFAULT)
	println("...done!\n")

	println(ctx.String())
}
