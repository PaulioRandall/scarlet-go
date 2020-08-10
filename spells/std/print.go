package std

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Print struct{}

func (Print) Summary() string {
	return `@Print(value...)
	Prints all arguments to standard output in the order provided.`
}

func (Print) Invoke(_ spellbook.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

type Println struct{}

func (Println) Summary() string {
	return `@Println(value...)
	Prints all arguments to standard output in the order provided then appends
	a linefeed.`
}

func (Println) Invoke(_ spellbook.Enviro, args []types.Value) {
	Print{}.Invoke(nil, args)
	fmt.Println()
}

func printSpellDocs() string {
	return `@Print(value...)    Prints all arguments to standard output in the order
	                  provided.
@Println(value...)  Same as @Print but appends a linefeed after the values.

Examples:

	# Outputs: "Hello, Scarlet!"
	@Print("Hello, Scarlet!")
	@Println("Hello, Scarlet!")

	# Outputs: "a*b = c"
	@Print(a, "*", b, " = ", c)
	@Println(a, "*", b, " = ", c)`
}
