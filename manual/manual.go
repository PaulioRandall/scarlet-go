package manual

import (
	"fmt"
	"strings"
)

func init() {
	Register("", overview)
}

type PageGenerator func() string

var generators = map[string]PageGenerator{}

func Register(pageName string, gen PageGenerator) {

	if gen == nil {
		panic(fmt.Errorf("Attempted to register nil page generator %q", pageName))
	}

	name := strings.ToLower(pageName)

	if _, ok := generators[name]; ok {
		panic(fmt.Errorf("Page generator with name %q already registered", pageName))
	}

	generators[name] = gen
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
	return `
Scarlet is a template for building domain or team specific scripting tools.

	"Sometimes it's better to light a flamethrower than curse the darkness."
		- 'Men at Arms' by Terry Pratchett

Usage:

	scarlet docs [search term]

Search terms:

	manifesto              Concepts & principles
	how                    The base language syntax & rules
	spells                 Available spells
	@<spell_name>          Specific spell documentation
	example                An example scroll
	versions               List of versions and their changes
	future                 Expected future changes`
}
