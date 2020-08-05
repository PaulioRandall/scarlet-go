package docs

import (
	"fmt"
	"strings"
)

type PageGenerator func() string

var generators = map[string]PageGenerator{}

func Register(pageName string, gen PageGenerator) {

	if pageName == "" {
		panic(fmt.Errorf("Attempted to register page with no name"))
	}

	if gen == nil {
		panic(fmt.Errorf("Attempted to register nil page generator %q", pageName))
	}

	name := strings.ToLower(pageName)

	if _, ok := generators[name]; ok {
		panic(fmt.Errorf("Page generator with name %q already registered", pageName))
	}

	generators[name] = gen
}

func Docs(searchTerm string) (int, error) {

	if strings.HasPrefix(searchTerm, "@") {
		return searchSpellDocs(searchTerm)
	}

	if searchTerm == "" {
		printOverview()
		return 0, nil
	}

	term := strings.ToLower(searchTerm)
	gen := generators[term]

	if gen == nil {
		return 0, fmt.Errorf("No documentation for %q", searchTerm)
	}

	page := gen()
	fmt.Println(page)
	return 0, nil
}

func searchSpellDocs(searchTerm string) (int, error) {
	return 0, fmt.Errorf("Spell documentation is not yet supported")
}
