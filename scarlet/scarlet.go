package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"

	"github.com/PaulioRandall/scarlet-go/pkg/gytha/parser"
	"github.com/PaulioRandall/scarlet-go/pkg/gytha/runtime"
	"github.com/PaulioRandall/scarlet-go/pkg/gytha/sanitiser"
	"github.com/PaulioRandall/scarlet-go/pkg/gytha/scanner"
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

	println("# Sanitising...")
	tks = sanitiser.SanitiseAll(tks)

	println("# Parsing...")
	stats := parser.ParseAll(tks)

	println("# Executing...")
	ctx := runtime.Run(stats)

	println()
	println(ctx.String())
}
