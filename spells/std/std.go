package std

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func init() {
	manual.Register("@exit", exitSpellDocs)
}

func Default() {
	InscribeAll(func(name string, spell spellbook.Spell) {
		e := spellbook.Inscribe(""+name, spell)
		if e != nil {
			panic(e)
		}
	})
}

func InscribeAll(inscribe spellbook.Inscriber) {
	inscribe("exit", Exit{})
	inscribe("print", Print{})
	inscribe("println", Println{})
	inscribe("set", Set{})
	inscribe("del", Del{})
}

type Exit struct{}

func (Exit) Summary() string {
	return `@Exit(exitcode)
	Exit terminates the current script with a specific exit code.`
}

func exitSpellDocs() string {
	return `@Exit(exitcode)
	Exit terminates the current script with a specific exit code.

Examples:

	@Exit(0)
	@Exit(1)`
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
