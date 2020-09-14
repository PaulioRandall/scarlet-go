package spells

import (
	"sort"
	"strings"

	"github.com/PaulioRandall/scarlet-go/manual"
	"github.com/PaulioRandall/scarlet-go/spells/fmtr"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
	"github.com/PaulioRandall/scarlet-go/spells/std"
)

func NewSpellbook() spellbook.Spellbook {

	spbk := spellbook.Spellbook{}

	std.RegisterAll(spbk)
	fmtr.RegisterAll(spbk)

	registerSpellDocs(spbk)
	return spbk
}

func registerSpellDocs(spbk spellbook.Spellbook) {
	manual.Register("spells", func() string {

		names := spbk.Names()
		sort.Strings(names)

		sb := strings.Builder{}

		for i, v := range names {

			sp, _ := spbk.LookUp(v)

			if i != 0 {
				sb.WriteString("\n\n")
			}

			sb.WriteString(sp.Sig)
			sb.WriteString("\n\t")
			sb.WriteString(sp.Desc)
		}

		return sb.String()
	})
}
