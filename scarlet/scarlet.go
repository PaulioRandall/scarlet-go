package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/scarlet-go/pkg/program"
)

func main() {

	args := os.Args

	if len(args) == 1 {
		// Dev run `./godo run`
		e := program.ProcessCommand("run", []string{"test.scarlet"})
		checkError(e)

		todo()
		return
	}

	if len(args) < 2 {
		e := fmt.Errorf("Missing command!")
		program.NewGenErr(e)
		checkError(e)
	}

	e := program.ProcessCommand(args[1], args[2:])
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
	println("[Next] Create a build error")
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
