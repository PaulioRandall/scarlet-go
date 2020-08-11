package std

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Set struct{}

func (Set) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 2 {
		env.Fail(errors.New("@Set requires two arguments"))
		return
	}

	idStr, ok := args[0].(types.Str)
	id := string(idStr)

	if !ok || !isIdentifier(id) {
		env.Fail(errors.New("@Set requires the first argument be an identifier string"))
		return
	}

	env.Bind(id, args[1])
}

func (Set) Docs() spellbook.SpellDoc {
	return man_setSpell
}

var man_setSpell = spellbook.SpellDoc{
	Pattern: `@Set("identifier", value)`,
	Summary: `Sets the value of variable represented by the first argument as the second
argument.`,
	Examples: []string{
		`@Set("x", 1)`,
		`@Set("name", "Scarlet")`,
	},
}

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
