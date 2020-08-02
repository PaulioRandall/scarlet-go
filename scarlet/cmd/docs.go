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
	<spell-name>       Documentation for a specific spell

Scarlet:

	Scarlet is a template for building domain specific scripting tools such as
	a replacement for bash. First I will present the principles that guide
	development as this should provide a good feel for why Scarlet was built and
	were its intended uses are:

	1. No script dependencies

		Scarlett scripts have no intrinsic way to import other Scarlett scripts to
		avoid the	many considerations and issues associated with the practice.
		Scarlet can get away with this as its designed specifcally to solve small
		problems such as those too big for bash but too small to be done nicely in
		more comprehensive languages, i.e C, Go, Rust, Java, etc.

	2. Build your own

		Scarlet emphasises the creation of spells (inbuilt functions) and
		recompilation rather than	importing libraries with the tool being designed
		to make this as easy as possible. Spells can be developed for your
		particular domain using the patterns you see fit rather than forcing a
		certain standard. The potential drawback is the lack of uniformity across
		domain specific tools and scripts derived from Scarlet, however, I would
		argue uniformaity is not an essential characteristic in this instance.

	3. Platform specific scripts

		TODO

	4. Minimalism, favour spells over native syntax

		TODO

	5. Don't like it, change it

		TODO

	6. Embed in your repository

		TODO

	7. Discworld Themed

		TODO

	Use cases driving development:

		TODO: Embed in repository to perform language independent API testing
		      without heavy testing tools

		TODO: Embed in repository to perform general configuration and deployment
		      activities

		TODO: Building programs for other languages

		TODO: General scripting of small problems such as iterating files and
		      data file transformations
`

	fmt.Println(s)
}
