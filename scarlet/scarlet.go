package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/scarlet/cmd"
)

func main() {

	if len(os.Args) == 1 {
		// Dev run `./godo run`
		args := cmd.NewArgs([]string{"run", "test.scroll"})
		eskarina(args)
		return
	}

	args := cmd.NewArgs(os.Args[1:])
	eskarina(args)
}

func eskarina(args cmd.Arguments) {

	exitCode, e := cmd.Run(args)
	if e != nil {
		fmt.Printf("[ERROR] %d\n%s\n", 1, e.Error())
	}

	os.Exit(exitCode)
}
