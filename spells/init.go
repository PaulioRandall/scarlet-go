package spells

import (
	"github.com/PaulioRandall/scarlet-go/spells/fmtr"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/std"
)

func init() {
	ns := namespacer("")
	std.InscribeAll(ns)
	fmtr.InscribeAll(ns)
}

func namespacer(prefix string) spellbook.Inscriber {
	return func(name string, spell spellbook.Spell) {
		e := spellbook.Inscribe(prefix+name, spell)
		if e != nil {
			panic(e)
		}
	}
}

func LookUp(name string) spellbook.Spell {
	return spellbook.LookUp(name)
}

func Inscribe(name string, spell spellbook.Spell) error {
	return spellbook.Inscribe(name, spell)
}

func ScribbleOut(name string) error {
	return spellbook.ScribbleOut(name)
}
