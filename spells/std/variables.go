package std

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Del struct{}

func (Del) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("@Del requires one argument"))
		return
	}

	id, ok := args[0].(types.Str)
	if !ok {
		env.Fail(errors.New("@Del requires its argument be an identifier string"))
		return
	}

	env.Unbind(string(id))
}

func (Del) Docs() spellbook.SpellDoc {
	return man_delSpell
}

var man_delSpell = spellbook.SpellDoc{
	Pattern: `@Del("identifier")`,
	Summary: `Deletes the variable represented by the first argument.`,
	Examples: []string{
		`@Del("x")`,
		`@Del("name")`,
	},
}

func isIdentifier(id string) bool {

	for i, ru := range id {

		if i == 0 {
			if !unicode.IsLetter(ru) {
				return false
			}

			continue
		}

		if !unicode.IsLetter(ru) || ru != '_' {
			return false
		}
	}

	return true
}
