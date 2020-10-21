package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/scarlet/cmd"
	"github.com/PaulioRandall/scarlet-go/scarlet/cmd2"
)

func main() {

	var e error
	a := cmd2.NewArgs(os.Args[1:])
	c, e := cmd2.Capture(a)
	checkErr(e, 1)

	switch v := c.(type) {
	case cmd2.HelpCmd:
		cmd2.Help(v)

	case cmd2.BuildCmd:
		_, e := cmd2.Build(v)
		checkErr(e, 1)

	case cmd2.RunCmd:
		r, e := cmd2.Run(v)
		checkErr(e, 1)
		if !r.Ok() {
			checkErr(r, r.ExitCode())
		}
		os.Exit(r.ExitCode())
	}
}

func checkErr(e error, errCode int) {
	if e != nil {
		fmt.Printf("[ERROR] %d\n%s\n", errCode, e.Error())
		os.Exit(errCode)
	}
}

func main_old() {
	args := cmd.NewArgs(os.Args[1:])

	exitCode, e := cmd.Run(args)

	// TODO: Exploit new error interface
	if e != nil {
		fmt.Printf("[ERROR] %d\n%s\n", 1, e.Error())
	}

	os.Exit(exitCode)
}
