package main

import (
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/program"
)

func main() { // Run it with `./godo run`
	e := program.Begin(os.Args)
	if e != nil {
		panic(e)
	}
}
