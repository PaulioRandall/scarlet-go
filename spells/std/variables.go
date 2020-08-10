package std

import (
	"errors"
	"unicode"

	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

type Set struct{}

func (Set) Summary() string {
	return `@Set("identifier", value)
	Sets the value of variable represented by the first argument as the second
	argument.`
}

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

type Del struct{}

func (Del) Summary() string {
	return `@Del("identifier")
	Deletes the variable represented by the first argument.`
}

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

func varSpellDocs() string {
	return `
@Set("identifier", value)
	Sets the value of variable represented by the first argument as the second
	argument.
@Del("identifier")
	Deletes the variable represented by the first argument

Examples:

	@Set("x", 1)
	@Set("name", "Scarlet")

	@Del("x")
	@Del("name")`
}
