package fmtr

import (
	"errors"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/types"
)

func InscribeAll(inscribe spellbook.Inscriber) {
	inscribe("FmtScript", Formatter{})
	manual.Register("@fmtscript", func() string {
		return spellbook.FmtSpellDoc(docs)
	})
}

type Formatter struct{}

var docs = spellbook.SpellDoc{
	Pattern: `@FmtScript(filename)`,
	Summary: `Attempts to format the specified Scarlett script.`,
	Examples: []string{
		`@FmtScript("./myscript.scar")`,
	},
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

func (Formatter) Docs() spellbook.SpellDoc {
	return docs
}
