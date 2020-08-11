package std

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Exit struct{}

var man_exitSpell = spellbook.SpellDoc{
	Pattern: `@Exit(exitCode)`,
	Summary: `Exit terminates the current script with a specific exit code.`,
	Examples: []string{
		"@Exit(0)",
		"@Exit(1)",
	},
}

func (Exit) Invoke(env spellbook.Enviro, args []types.Value) {

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

func (Exit) Docs() spellbook.SpellDoc {
	return man_exitSpell
}
