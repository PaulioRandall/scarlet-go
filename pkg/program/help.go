package program

import (
	"fmt"
)

// help show the instructions and options for how to use Scarlet CLI.
func help(args []string) error {

	size := len(args)

	switch {
	case size > 1:
		return fmt.Errorf("Unexpected argument %q", args[1])

	case size == 0:
		printHelp()

	case args[0] == "build":
		printBuildHelp()

	case args[0] == "run":
		printRunHelp()

	default:
		return fmt.Errorf("Unexpected argument %q", args[0])
	}

	return nil
}

var buildOpts = map[string]string{
	"nofmt": "Don't format the script.",
	"log":   "Logs the output of each stage into labelled files within the specified folder.",
}

func printHelp() {

	s := `Scarlet is a tool for parsing and executing Scarlett scripts.

Usage:

	scarlet <command> [arguments]

Commands:

	help [<command>]      Show CLI instructions.
	docs [<search term>]  Show language documentation.
	build                 Parses, compiles, and formats the script.
	run                   Parses, compiles, formats, then executes the script.
`

	println(s)
}

func printBuildHelp() {

	s := `'build' compiles and validates a script.

Usage:

	scarlet build [options] <script file>

Options:

	-nofmt
		Don't format the script.
	-log <output folder>
		Logs the output of each compilation stage as labelled files into the
		output folder.
`

	println(s)
}

func printRunHelp() {

	s := `'run' compiles, validates, then runs a script.

Usage:

	scarlet run [options] <script file>

Options:

	-nofmt
		Don't format the script.
	-log 'output folder'
		Logs the output of each compilation stage as labelled files into the
		'output folder'.
`

	println(s)
}
