package cmd

import (
	"fmt"
)

// Help show the instructions and options for how to use Scarlet CLI.
func Help(args Arguments) (int, error) {

	switch {
	case args.Count() > 1:
		args.Shift()
		return 1, fmt.Errorf("Unexpected argument %q", args.Peek())

	case args.Empty():
		printHelp()

	case args.Peek() == "build":
		printBuildHelp()

	case args.Peek() == "run":
		printRunHelp()

	default:
		return 1, fmt.Errorf("Unexpected argument %q", args.Peek())
	}

	return 0, nil
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

	fmt.Println(s)
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

	fmt.Println(s)
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

	fmt.Println(s)
}
