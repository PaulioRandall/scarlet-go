package std

import (
	"errors"
	"fmt"

	//"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func RegisterAll() {

	spellbook.Register(Spell_Exit, spellbook.SpellDoc{
		Name: "Exit",
		Sig:  "@Exit(exitCode)",
		Desc: "Exit terminates the current scroll with a specific exit code.",
		Examples: []string{
			"@Exit(0)",
			"@Exit(1)",
		},
	})

	spellbook.Register(Spell_Print, spellbook.SpellDoc{
		Name: "Print",
		Sig:  "@Print(value...)",
		Desc: "Prints all arguments, in the order provided, to standard output.",
		Examples: []string{
			`@Print("Hello, Scarlet!")`,
			`@Print(a, "*", b, " = ", c)`,
		},
	})

	spellbook.Register(Spell_Println, spellbook.SpellDoc{
		Name: "Println",
		Sig:  "@Println(value...)",
		Desc: "Same as @Print but appends a linefeed.",
		Examples: []string{
			`@Println("Hello, Scarlet!")`,
			`@Println(a, "*", b, " = ", c)`,
		},
	})
}

func Spell_Exit(env spellbook.Enviro, args []types.Value) {

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

func Spell_Print(_ spellbook.Enviro, args []types.Value) {
	for _, v := range args {
		fmt.Print(v.String())
	}
}

func Spell_Println(_ spellbook.Enviro, args []types.Value) {
	Spell_Print(nil, args)
	fmt.Println()
}
