package spells

import (
	"sort"
	"strings"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/fmtr"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/std"
)

func init() {
	ns := namespace("")
	std.InscribeAll(ns)
	fmtr.InscribeAll(ns)
	manual.Register("spells", spellDocs)
}

func namespace(prefix string) spellbook.Inscriber {
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

func spellDocs() string {

	names := spellbook.SpellNames()
	sort.Strings(names)

	sb := strings.Builder{}

	for i, v := range names {

		sp := spellbook.LookUp(v)

		if i != 0 {
			sb.WriteString("\n\n")
		}

		sb.WriteString(sp.Summary())
	}

	return sb.String()
}
