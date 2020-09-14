package manual

import (
	"sort"
	"strings"

	"github.com/PaulioRandall/scarlet-go/spells"
)

func spellsDoc() string {

	spbk := spells.NewSpellbook()
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
}
