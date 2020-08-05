package docs

import (
	"fmt"
	"sort"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
)

func printSpellOverview() {

	s := `Spells are the central concept on which Scarlet was built. A less
glamorous name would be 'inbuilt functions', but since a writer of Scarlett
scripts only needs to know how to use them, not how they work, I felt 'spells'
were a more fitting and engaging title.

	"Any sufficiently advanced technology is indistinguishable from magic"
		- [One of] 'Clarke's three laws' by Arthur C Clarke

Usage:

	@spell_name([argument, argument, ...])

Examples:

	@Set(x, 42)
	@Println("6 * 7 = ", x)
	@Exit(0)

Description:

	Spells are always prefixed with an at sign '@' followed by their name and
	accept arguments in the same manner as iconic C-style functions. Unlike
	variable names, spell names may contain dots '.' to mimic	namespaces. This
	can make spells more readable, better convey their usage,	and are easier
	to mass search-and-replace. A registered spell name may have as many
	namespace segments as the coder likes but they should strive to create
	names that are short and meaningful.

		@list.Add(x)
		x := @list.num.Sum()

Pros & cons:

	+ Can be high performant due to their Go implementations.
	+ Are more robust due to Go's stricter typing and safety features.
	+ Have access to Go's rapidly growing open source libraries.
	+ Ready to use without importing.
	+ Can be created, updated, or removed by any inquisitive programmer.
	+ Can be optimised for a very precise problem, domain, or environment. 
	+ Can easily do things that are tedious within Scarlett scripts. 

	- Require a knowledge of Go to create and modify.
	- Require Scarlet to be recompiled to be usable.

Future changes:

	I hope to add blocks as parameters to enable code to be run with the
	callers scope and variables. Here is an example spell with block parameter:

	@If(true, {
		@Set(x, 1)
	})

	In future I hope to add some very common but completely removable and
	modifable spells to get users started. However, many of these depend on
	language features not yet implemented:

		x := @Args()                    Get the program arguments
		x := @Exists("variable_name")   Does a variable exist?
		x := @Len(value)                Find the length of lengthed value
		x := @Str(value)                Stringify a value of any type
		@Panic(exitCode, message)       Exit the script with an error message
		e := @Catch({                   Catch any panics and return as an error
			...
		})

		@str.                           'List' type & spells
		@map.                           'Map' type & spells
		@fmt.                           'Template' type & spells
		@io.                            Basic input and output spells
`

	fmt.Println(s)
}

func printSpells() {

	names := spellbook.SpellNames()
	sort.Strings(names)

	for i, v := range names {

		sp := spellbook.LookUp(v)

		if i != 0 {
			fmt.Println()
		}

		fmt.Println(sp.Summary())
	}
}
