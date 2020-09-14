package manual

import (
	"fmt"
	"sort"
	"strings"

	"github.com/PaulioRandall/scarlet-go/spells"
	"github.com/PaulioRandall/scarlet-go/spells/spellbook"
)

func init() {
	Register("", overview)
	registerSpellDocs(spells.NewSpellbook())
}

func registerSpellDocs(spbk spellbook.Spellbook) {
	Register("spells", func() string {

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

type PageGenerator func() string

var generators = map[string]PageGenerator{}

func Register(pageName string, gen PageGenerator) {

	if gen == nil {
		panic(fmt.Errorf("Attempted to register nil page generator %q", pageName))
	}

	name := strings.ToLower(pageName)

	if pgExists(name) {
		panic(fmt.Errorf("Page generator with name %q already registered", pageName))
	}

	generators[name] = gen
}

func pgExists(name string) bool {
	pg, ok := generators[name]
	return ok && pg != nil
}

func LookUp(pageName string) PageGenerator {
	name := strings.ToLower(pageName)
	return generators[name]
}

func Search(searchTerm string) (string, bool) {

	term := strings.ToLower(searchTerm)
	gen := generators[term]

	if gen == nil {
		return "", false
	}

	return gen(), true
}

func overview() string {
	return `Scarlet is a template for building domain or team specific scripting tools.

	"Sometimes it's better to light a flamethrower than curse the darkness."
		- 'Men at Arms' by Terry Pratchett

Usage:

	scarlet docs [search term]

Search terms:

	manifesto              Concepts & principles
	syntax | rules         The base language syntax & rules
	spells                 Available spells
	examples               An example scroll
	chapters               List of chapters and changes introduced
	future                 Potential future changes
	@<spell_name>          Specific spell documentation`
}
