package program

import (
	"fmt"
)

// help show the instructions and options for how to use Scarlet CLI.
func help(args []string) error {

	size := len(args)
	if size > 1 {
		return fmt.Errorf("Unexpected argument %q", args[0])
	}

	if size == 0 {
		printHelp()
	} else if args[0] == "docs" {
		printDocsHelp()
	} else if args[0] == "build" {
		printBuildHelp()
	} else if args[0] == "run" {
		printRunHelp()
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

func printDocsHelp() {

}

func printBuildHelp() {

}

func printRunHelp() {

}
