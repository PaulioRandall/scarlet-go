package std

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func InscribePrint(inscribe spellbook.Inscriber) {

	inscribe("print", Print{})
	inscribe("println", Println{})

	manual.Register("@print", printSpellDocs)
	manual.Register("@println", printSpellDocs)
}

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
	Prints all arguments to standard output in the order given, then appends
	a linefeed.`
}

func (Println) Invoke(_ spellbook.Enviro, args []types.Value) {
	Print{}.Invoke(nil, args)
	fmt.Println()
}

func printSpellDocs() string {
	return `
@Print(value, value, ...)
	Prints all arguments to standard output in the order given.

@Println(value, value, ...)
	Same as @Print but appends a linefeed.

Examples

	@Print("Hello, Scarlet!")
	@Print(a, "*", b, " = ", c)

	@Println("Hello, Scarlet!")
	@Println(a, "*", b, " = ", c)`
}
