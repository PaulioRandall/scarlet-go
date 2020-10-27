package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/cmd/program"
)

func main() {

	var e error
	a := program.NewArgs(os.Args[1:])
	c, e := program.Capture(a)
	checkErr(e, 1)

	switch v := c.(type) {
	case program.HelpCmd:
		program.Help(v)

	case program.BuildCmd:
		_, e := program.Build(v)
		checkErr(e, 1)

	case program.RunCmd:
		r, e := program.Run(v)
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
