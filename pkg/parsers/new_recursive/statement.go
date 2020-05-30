package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type parser struct {
	*pipeline
	Factory
}

func ParseStatements(fac Factory, tks []Token) ([]Statement, error) {

	var (
		p = &parser{newPipeline(tks), fac}
		s = []Statement{}
	)

	for p.hasMore() {

		expr, e := parseExpression(p)
		if e != nil {
			return nil, e
		}

		s = append(s, expr)

		_, e = p.expect(TERMINATOR)
		if e != nil {
			return nil, e
		}
	}

	return s, nil
}

func parseStatement(p *parser) (Statement, error) {
	return parseExpression(p)
}
