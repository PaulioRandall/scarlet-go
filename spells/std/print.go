package std

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Print struct{}

var man_printSpell = spellbook.SpellDoc{
	Pattern: `@Print(value, value, ...)`,
	Summary: `Prints all arguments to standard output in the order given.`,
	Examples: []string{
		`@Print("Hello, Scarlet!")`,
		`@Print(a, "*", b, " = ", c)`,
	},
}

func (Print) Summary(name string) string {
	return name + `(value...)
	Prints all arguments to standard output in the order provided.`
}

func (Print) Invoke(_ spellbook.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func (Print) Docs() spellbook.SpellDoc {
	return man_printSpell
}

type Println struct{}

var man_printlnSpell = spellbook.SpellDoc{
	Pattern: `@Println(value, value, ...)`,
	Summary: `Prints all arguments to standard output in the order given, then appends
a linefeed.`,
	Examples: []string{
		`@Println("Hello, Scarlet!")`,
		`@Println(a, "*", b, " = ", c)`,
	},
}

func (Println) Invoke(_ spellbook.Enviro, args []types.Value) {
	Print{}.Invoke(nil, args)
	fmt.Println()
}

func (Println) Docs() spellbook.SpellDoc {
	return man_printlnSpell
}
