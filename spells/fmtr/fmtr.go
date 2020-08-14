package fmtr

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func InscribeAll(inscribe spellbook.Inscriber) {
	inscribe("FmtScript", Formatter{})
	manual.Register("@fmtscript", fmtrSpellDocs)
}

type Formatter struct{}

func (Formatter) Summary() string {
	return `@FmtScript(filename)
	Attempts to format the specified Scarlett script.`
}

func (Formatter) Invoke(env spellbook.Enviro, args []types.Value) {

	if len(args) != 1 {
		env.Fail(errors.New("Formatting requires a single filename argument"))
		return
	}

	filename, ok := args[0].(types.Str)
	if !ok {
		env.Fail(errors.New("Formatting requires its argument be a filename"))
		return
	}

	e := formatFile(filename.String())
	if e != nil {
		env.Fail(e)
	}
}

func fmtrSpellDocs() string {
	return `
@FmtScript(filename)
	Attempts to format the specified Scarlett script.

Examples

	@FmtScript("./myscript.scar")`
}
