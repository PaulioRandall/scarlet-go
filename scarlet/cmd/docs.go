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

	Scarlet is a template for building domain or environment specific
	scripting	tools such as bash replacements. First I will present the
	desirable characteristics that guide development then a number of
	envisioned use cases. This should provide a good feel for why Scarlet
	was built	and	its intended purposes.

	1. No dependencies

		Scarlett scripts have no native way to import other Scarlett scripts.
		This avoids the	many considerations and issues associated with the
		practice. Scarlet is designed specifcally for small problems and
		networkless environments.

		"Darkness isn't the opposite of light, it is simply its absence."
			- 'The Light Fantastic' by Terry Pratchett

	2. Easy integration

		Scarlet emphasises the creation of spells (inbuilt functions) rather
		than importing libraries. Spells are written in Go so external Go
		libraries can be used. Simply register the spell and recompile Scarlet.
		I envisioned the tool will be copied and then populated with domain or
		environment specific spells using any patterns the authors see fit.

		"The three rules of the Librarians of Time and Space are:
		1) Silence;
		2) Books must be returned no later than the date last shown; and
		3) Do not interfere with the nature of causality."
			- 'Guards! Guards!' by Terry Pratchett

	3. Context specific

		Unlike some modern scripting languages, Scarlett scripts are designed
		to be platform and situation specific, that is, scripts are written for
		a single purpose and against a specific version of the tool. This may
		seem rather restrictive but its to encourage context driven solutions
		and surpress the compelling urge to abstract everything.

		"THAT'S MORTALS FOR YOU. THEY'VE ONLY GOT A FEW YEARS IN THIS WORLD AND
		THEY SPEND THEM ALL IN MAKING THINGS COMPLICATED FOR THEMSELVES."
			- 'Mort' by Terry Pratchett

	4. Minimalist

		Scarlet favours spells over native syntax, vis if a feature is not used
		constantly or is niche then its better of as a spell that can more
		easily be modified. Fewer default native features allows others to be
		added quickly when desired.

		"Take it from me, there's nothing more terrible than someone out to do
		the world a favour."
			- 'Sourcery' by Terry Pratchett

	5. Customisable

		If you don't like the names of spells, change them.
		If you don't like the language keywords, change them.
		If you don't like function brackets, change them.

		"What don’t die can’t live.
		What don’t live can’t change.
		What don’t change can’t learn."
			- 'Lords and Ladies' by Terry Pratchett

	6. Light and portable

		The Scarlet executable is light, portable, and does not require an
		installation process; much like Lua. With time and hope a Rust
		implementation will be built precisely for embedding in other programs
		and repositories.

		"'What's a philosopher?' said Brutha. 'Someone who's bright enough to
		find a job with no heavy lifting,' said a voice in his head."
			- 'Small Gods' by Terry Pratchett

	7. Discworld Themed

		Scarlet is a discworld themed tool, because the logical domain that is
		programming is filled with things that don't make sense unless you were
		in the right place, at the right time, inside the right head. There's
		also a servere lack of true magic, probably why ladies arn't that
		interested.

		"The Turtle Moves"
			- 'Small Gods' by Terry Pratchett

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
