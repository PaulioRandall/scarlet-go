package main

import (
	"io/ioutil"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/err"

	parser "github.com/PaulioRandall/scarlet-go/pkg/parsers/gytha"
	runtime "github.com/PaulioRandall/scarlet-go/pkg/runtime/cutangle"
	sanitiser "github.com/PaulioRandall/scarlet-go/pkg/sanitisers/standard"
	scanner "github.com/PaulioRandall/scarlet-go/pkg/scanners/lupine"
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
