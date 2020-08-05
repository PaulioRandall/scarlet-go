package manual

import (
	"fmt"
	"strings"
)

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
