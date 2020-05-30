package recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/parsers/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type parser struct {
	*pipeline
	Factory
}

func ParseAll(f Factory, tks []Token) ([]Expression, error) {

	p := &parser{
		newPipeline(tks),
		f,
	}

	return parseExpressions(p)
}
