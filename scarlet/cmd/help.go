package cmd

import (
	"fmt"
)

// Help show the instructions and options for how to use Scarlet CLI.
func help(args Arguments) (int, error) {

	switch {
	case args.count() > 1:
		args.shift()
		return 1, fmt.Errorf("Unexpected argument %q", args.peek())

	case args.empty():
		printHelp()

	case args.peek() == "build":
		printBuildHelp()

	case args.peek() == "run":
		printRunHelp()

	default:
		return 1, fmt.Errorf("Unexpected argument %q", args.peek())
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
		Don't format the script after building.
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
`

	fmt.Println(s)
}
