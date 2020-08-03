package docs

import (
	"fmt"
)

func printSpellOverview() {

	s := `Spells.

Usage:

	@spell_name([argument, argument, ...])

Examples:

	@Set(x, 42)
	@Println("Answer is ", x)
	@Exit(0)

Description:

	"Any sufficiently advanced technology is indistinguishable from magic"
		- 'Clarke's three laws' by Arthur C Clarke

	Spells are one of the core concepts on which Scarlet was built. A less
	glamorous description would be 'inbuilt functions', but since a writer
	of Scarlett scripts only needs to know how to use them, not how they work,
	I felt 'spells' were a more fitting and engaging title.

	They:
	+ Can be high performant due to their Go implementations.
	+ Are more robust due to Go's stricter static typing and safety features.
	+ Have access to Go's rapidly growing collection of open source libraries.
	+ Ready to use without importing.
	+ Can be created, updated, or removed by any inquisitive programmer.
	+ Can be optimised for a very precise problem, domain, or environment. 
	+ Can easily do things that are tedious within Scarlett scripts. 
	- Require a knowledge of Go to modify, not much of a negative in honesty.
	- Require Scarlet to be recompiled to be usable.

	Spells are always prefixed with an at sign '@' followed by their name.
	Unlike variable names, spell names may contain dots '.' to mimic
	namespaces,	called 'spellbooks'. This can make spells more readable,
	better convey their usage, and are easier to mass search-and-replace.

	Spells accept arguments in the same manner as classical functions and once
	implemented, will allow multiple returns too. I also hope to allow blocks
	as parameters to enable code to be run with the callers scope and
	variables. They might look something like this:

	@If(true, {
		@Set(x, 1)
	})

	In future I hope to add some very common but completely removable and
	modifable spells to get users started. However, many of these depend on
	language features not yet implemented:

		x := @Args()                  # Get the program arguments
		x := @Exists("variable_name") # Does a variable exist?
		x := @Len(value)              # Find the length of lengthed value
		x := @Str(value)              # Stringify a value of any type
		@Panic(exitCode, message)     # Exit the script with an error message
		e := @Catch({                 # Catch any panics and return as an error
			...
		})

		@str.    # 'List' type & spells
		@map.    # 'Map' type & spells
		@fmt.    # 'Template' type & spells
		@io.     # Basic input and output spells
`

	fmt.Println(s)
}
