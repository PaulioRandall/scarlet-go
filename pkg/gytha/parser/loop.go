package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
)

func isLoop(p *pipe) bool {
	return p.match(TK_LOOP)
}

func parseLoop(p *pipe) Statement {
	// pattern := LOOP ID guard
	// pattern := LOOP ID DELIM ID DELIM ID UPDATES expression

	key := p.expect(`parseLoop`, TK_LOOP)
	indexId := p.expect(`parseLoop`, TK_IDENTIFIER)

	if p.accept(TK_ASSIGNMENT) {

		initIndex := parseExpression(p)

		if initIndex == nil {
			err.Panic(
				errMsg("parseLoop", `expression`, p.peek()),
				err.At(p.peek()),
			)
		}

		return Loop{
			Open:      key,
			IndexId:   indexId,
			InitIndex: initIndex,
			Guard:     parseGuard(p),
		}
	}

	p.expect(`parseLoop`, TK_DELIMITER)
	valueId := p.expect(`parseLoop`, TK_IDENTIFIER)

	p.expect(`parseLoop`, TK_DELIMITER)
	moreId := p.expect(`parseLoop`, TK_IDENTIFIER)

	p.expect(`parseLoop`, TK_UPDATES)

	return ForEach{
		Open:    key,
		IndexId: indexId,
		ValueId: valueId,
		MoreId:  moreId,
		List:    parseExpression(p),
		Block:   parseBlock(p),
	}
}
