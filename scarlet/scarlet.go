package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/program"
)

func main() {

	if len(os.Args) == 1 {
		// Dev run `./godo run`
		args := program.NewArgs([]string{"run", "test.scarlet"})
		e := program.Execute(args)
		checkError(e)

		todo()
		return
	}

	args := program.NewArgs(os.Args[1:])
	e := program.Execute(args)
	checkError(e)
}

func checkError(e error) {

	if e == nil {
		return
	}

	fmt.Println(e.Error())

	if se, ok := e.(program.ScarletError); ok {
		os.Exit(se.ExitCode)
	}

	os.Exit(1)
}

func todo() {
	println()
	println()
	println("TODO:")
	println("[Next] Add compile stage output logging")
	println("[Next] Create formatting tool")
	println("[Next] Add scanning for complex expressions")
	println()
	println("[Plan]")
	println("- a_scan:     scans in tokens including redundant ones")
	println("- b_sanitise: removes redundant tokens")
	println("- c_check:    checks the token sequence follows language rules")
	println("- d_shunt:    converts from infix to postfix notation")
	println("- e_compile:  converts the tokens into instructions")
	println("- f_runtime:  executes an instruction list")
	println("- ...")
}
