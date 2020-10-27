package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/scarlet/cmd"
)

func main() {

	var e error
	a := cmd.NewArgs(os.Args[1:])
	c, e := cmd.Capture(a)
	checkErr(e, 1)

	switch v := c.(type) {
	case cmd.HelpCmd:
		cmd.Help(v)

	case cmd.BuildCmd:
		_, e := cmd.Build(v)
		checkErr(e, 1)

	case cmd.RunCmd:
		r, e := cmd.Run(v)
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
