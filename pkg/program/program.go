package program

import (
	"fmt"
)

// scarlet command [options] script
//
// Commands:
//   help 								Show the instructions and options for how to use
//                    		Scarlet CLI
//   docs [search term]		Show documentation.
//   docs help						Show the instructions and options for navigating
//            						the documentation.
//   build								Parses, compiles, and formats the script.
//   build help						Show the instructions and options for building.
//   run  								Parses, compiles, formats, and runs the script.
//   run help							Show the instructions and options for running
//
// Options:
//   -nofmt								Don't format the script.
//   -log folder					Logs the output of each stage into labelled files
//              					within the specified folder.

func Begin(args []string) error {

	if len(args) < 2 {
		return fmt.Errorf("Missing command!")
	}

	return processCmd(args[1], args[2:])
}

func processCmd(cmd string, args []string) error {

	switch cmd {
	case "todo":
		todo()
		return nil

	case "help":
		return help(args)

	case "docs":
		return docs(args)

	case "build":
		_, e := build(args)
		return e

	case "run":
		return run(args)
	}

	return fmt.Errorf("Unknown command %q", cmd)
}

func todo() {
	println()
	println()
	println("TODO:")
	println("[Next] Implement help command")
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
