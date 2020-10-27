package cmd

import (
	"fmt"
)

const usage = `Scarlet is a scripting tool for parsing and executing Scarlet scrolls.

Usage:

	scarlet <command> [arguments]

Commands:

	help [command]         Show CLI instructions.
	build                  Parses and compiles a scroll.
	run                    Parses, compiles, then executes a scroll.
`

// Help displays Scarlet CLI usage.
func Help(c HelpCmd) {
	fmt.Println(usage)
}
