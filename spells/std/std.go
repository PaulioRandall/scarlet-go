package std

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func RegisterAll(sb spellbook.Spellbook) error {

	sb.Register(
		Spell_Exit,
		"Exit",
		"@Exit(exitCode)",
		"Exit terminates the current scroll with a specific exit code.",
		"@Exit(0)",
		"@Exit(1)",
	)

	sb.Register(
		Spell_Print,
		"Print",
		"@Print(value...)",
		"Prints all arguments, in the order provided, to standard output.",
		`@Print("Hello, Scarlet!")`,
		`@Print(a, "*", b, " = ", c)`,
	)

	sb.Register(
		Spell_Println,
		"Println",
		"@Println(value...)",
		"Same as @Print but appends a linefeed.",
		`@Println("Hello, Scarlet!")`,
		`@Println(a, "*", b, " = ", c)`,
	)

	return nil
}

func Spell_Exit(_ spellbook.Entry, env spellbook.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Exit requires one argument"))
		return
	}

	if c, ok := args[0].(types.Num); ok {
		env.Exit(int(c.Integer()))
		return
	}

	env.Fail(errors.New("@Exit requires its argument be a number"))
}

func Spell_Print(_ spellbook.Entry, _ spellbook.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func Spell_Println(spell spellbook.Entry, _ spellbook.Enviro, args []types.Value) {
	Spell_Print(spell, nil, args)
	fmt.Println()
}
