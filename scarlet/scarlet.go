package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/scarlet/cmd"
)

func main() {
	args := cmd.NewArgs(os.Args[1:])

	exitCode, e := cmd.Run(args)

	// TODO: Exploit new error interface
	if e != nil {
		fmt.Printf("[ERROR] %d\n%s\n", 1, e.Error())
	}

	os.Exit(exitCode)
}
