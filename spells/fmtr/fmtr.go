package fmtr

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func InscribeAll(inscribe spellbook.Inscriber) {
	inscribe("FmtScroll", Formatter{})
	manual.Register("@fmtscroll", fmtrSpellDocs)
}

type Formatter struct{}

func (Formatter) Summary() string {
	return `@FmtScroll(filename)
	Attempts to format the specified Scarlet scroll.`
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
@FmtScroll(filename)
	Attempts to format the specified Scarlet scroll.

Examples

	@FmtScroll("./example.scroll")`
}
