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
	<spell-name>       Documentation for a specific spell, e.g. '@Set'

Scarlet:

	Scarlet is a template for building domain or environment specific scripting
	tools such as bash replacements. First I will present the principles that
	guide development then a number of envisioned use cases as this should provide
	a good feel for why Scarlet was built	and	its intended purpose:

	1. No script dependencies

		Scarlett scripts have no intrinsic way to import other Scarlett scripts to
		avoid the	many considerations and issues associated with the practice.
		Scarlet is designed specifcally for small problems and networkless
		environments.

	2. Build your own

		Scarlet emphasises the creation of spells (inbuilt functions) rather than
		importing libraries. Spells are written in Go so external Go libraries can
		be used. Simple register the spell and recompile Scarlet.	I envisioned the
		tool will be copied and then populated with domain or	environment specific
		spells using any patterns the authors see fit.

	3. Context specific

		Unlike other modern scripting languages, Scarlett scripts are designed to be
		platform and situation specific, that is, scripts are written for a single
		purpose and against a specific version of the tool. This may seem rather
		restrictive but its to encourage context driven solutions and surpress the
		compelling urge to abstract everything.

	4. Minimalism

		Scarlet favours spells over native syntax, vis if a feature is not used
		constantly or is niche then its better of as a spell that can more easily
		be modified. Fewer default native features allows others to be added
		quickly when desired.  

	5. Don't like it, change it

		If you don't like the names of spells, change them.
		If you don't like the language keywords, change them.
		If you don't like parenthesis on functions, change them.

		"What don’t die can’t live.
			What don’t live can’t change.
				What don’t change can’t learn."
					- Lords and Ladies, Terry Pratchett

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
