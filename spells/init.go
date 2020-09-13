package spells

import (
	"github.com/PaulioRandall/scarlet-go/spells/fmtr"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/std"
)

func init() {
	std.RegisterAll()
	fmtr.RegisterAll()
}

func LookUp(name string) (spellbook.Entry, bool) {
	sp, ok := spellbook.LookUp(name)
	return sp, ok
}
