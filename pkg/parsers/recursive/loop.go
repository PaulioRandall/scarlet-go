package recursive

import (
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

	if isGuard(p) {
		return Loop{
			Open:     key,
			IndexVar: indexId,
			Guard:    parseGuard(p),
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
