package cmd

import (
	"fmt"
)

//   docs [search term]		Show documentation.

func docs(args Arguments) (int, error) {

	switch {
	case args.empty():
		printDocs()
		return 0, nil

	default:
		return 1, fmt.Errorf("Unexpected argument %q", args.peek())
	}

	return 0, nil
}

func printDocs() {

	s := `Scarlet's language documentation.

Usage:

	scarlet docs [search term]

Search terms:

	spell              Spells and how to use them
	variable           How to use variables
	comments           Writing comments
	types              Variable types and their uses

Scarlet:

	TODO: Here will be a description of Scarlett's general language rules
	      and an overview of features.
`

	fmt.Println(s)
}
