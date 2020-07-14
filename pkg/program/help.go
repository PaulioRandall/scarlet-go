package program

import (
	"fmt"
)

// help show the instructions and options for how to use Scarlet CLI.
func help(args []string) error {

	size := len(args)

	switch {
	case size > 1:
		e := fmt.Errorf("Unexpected argument %q", args[1])
		return NewGenErr(e)

	case size == 0:
		printHelp()

	case args[0] == "build":
		printBuildHelp()

	case args[0] == "run":
		printRunHelp()

	default:
		e := fmt.Errorf("Unexpected argument %q", args[0])
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
