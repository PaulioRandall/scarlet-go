package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isLoop(p *pipe) bool {
	return p.match(LOOP)
}

func parseLoop(p *pipe) Statement {
	// pattern := LOOP ID guard
	// pattern := LOOP ID DELIM ID DELIM ID UPDATES expression

	key := p.expect(`parseLoop`, LOOP)
	indexId := p.expect(`parseLoop`, IDENTIFIER)

	if p.accept(ASSIGN) {

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

	p.expect(`parseLoop`, DELIMITER)
	valueId := p.expect(`parseLoop`, IDENTIFIER)

	p.expect(`parseLoop`, DELIMITER)
	moreId := p.expect(`parseLoop`, IDENTIFIER)

	p.expect(`parseLoop`, UPDATES)

	return ForEach{
		Open:    key,
		IndexId: indexId,
		ValueId: valueId,
		MoreId:  moreId,
		List:    parseExpression(p),
		Block:   parseBlock(p),
	}
}
