package program

import (
	"fmt"
)

// help show the instructions and options for how to use Scarlet CLI.
func help(args Arguments) error {

	switch {
	case args.count() > 1:
		args.take()
		e := fmt.Errorf("Unexpected argument %q", args.peek())
		return NewGenErr(e)

	case args.empty():
		printHelp()

	case args.peek() == "build":
		printBuildHelp()

	case args.peek() == "run":
		printRunHelp()

	default:
		e := fmt.Errorf("Unexpected argument %q", args.peek())
		return NewGenErr(e)
	}

	return nil
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
