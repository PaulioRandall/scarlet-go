package spellbook

import (
	"sort"
	"strings"

	"github.com/PaulioRandall/scarlet-go/manual"
)

func init() {
	manual.Register("spells", spellSummaries)
}

func spellSummaries() string {

	names := SpellNames()
	sort.Strings(names)

	sb := strings.Builder{}

	for i, v := range names {

		sp := LookUp(v)

		if i != 0 {
			sb.WriteString("\n\n")
		}

		s := sp.Summary()
		sb.WriteString(s)
	}

	return sb.String()
}
