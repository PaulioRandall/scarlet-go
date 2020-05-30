package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type parser struct {
	*pipeline
	Factory
}

func ParseStatements(fac Factory, tks []Token) ([]Statement, error) {

	r := []Statement{}
	p := &parser{newPipeline(tks), fac}

	for p.hasMore() {

		st, e := expectStatement(p)
		if e != nil {
			return nil, e
		}

		r = append(r, st)

		_, e = p.expect(TERMINATOR)
		if e != nil {
			return nil, e
		}
	}

	return r, nil
}

func parseStatement(p *parser) (Statement, error) {
	return parseExpression(p)
}

func expectStatement(p *parser) (Statement, error) {

	s, e := parseStatement(p)

	if e == nil && s == nil {
		return nil, err.New("Expected statement", err.At(p.any()))
	}

	return s, e
}
