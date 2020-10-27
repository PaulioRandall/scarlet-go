package manual

import (
	"fmt"
	"strings"
)

type PageGenerator func() string
type Manual map[string]PageGenerator

func New() Manual {

	m := Manual{}

	m.Register("", overview)
	m.Register("manifesto", manifesto)
	m.Register("syntax", syntax)
	m.Register("examples", examples)
	m.Register("future", future)
	m.Register("chapters", chapters)

	return m
}

func (m Manual) Register(name string, gen PageGenerator) {

	if gen == nil {
		panic(fmt.Errorf("Attempted to register nil page generator %q", name))
	}

	regName := strings.ToLower(name)

	if _, exists := m[regName]; exists {
		panic(fmt.Errorf("Page generator with name %q already registered", name))
	}

	m[regName] = gen
}

func (m Manual) LookUp(name string) PageGenerator {
	regName := strings.ToLower(name)
	return m[regName]
}

func (m Manual) Search(searchTerm string) (string, bool) {

	term := strings.ToLower(searchTerm)
	gen := m[term]

	if gen == nil {
		return "", false
	}

	return gen(), true
}
