package cmd

import (
	"fmt"
)

// Help show the instructions and options for how to use Scarlet CLI.
func help(args Args) (int, error) {

	switch {
	case args.count() > 1:
		args.shift()
		return 1, fmt.Errorf("Unexpected argument %q", args.peek())

	case args.empty():
		printHelp()

	case args.accept("build"):
		printBuildHelp()

	case args.accept("run"):
		printRunHelp()

	default:
		return 1, fmt.Errorf("Unexpected argument %q", args.peek())
	}

	return 0, nil
}

func printHelp() {

	s := `Scarlet is a tool for parsing and executing Scarlet scrolls.

Usage:

	scarlet <command> [arguments]

Commands:

	help [command]         Show CLI instructions.
	docs|man [search term] Show language documentation.
	build                  Parses and compiles a scroll.
	run                    Parses, compiles, then executes a scroll.
`

	fmt.Println(s)
}

func printBuildHelp() {

	s := `'build' compiles and validates a scroll.

Usage:

	scarlet build [options] <scroll file>

Options:

	-log <output directory>
		Log output of each parsing stage to a file.
`

	fmt.Println(s)
}

func printRunHelp() {

	s := `'run' compiles, validates, then runs a scroll.

Usage:

	scarlet run [options] <scroll file>

Options:

	-log <output directory>
		Log output of each parsing stage to a file.
`

	fmt.Println(s)
}
