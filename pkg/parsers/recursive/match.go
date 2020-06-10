package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/err"
	. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isWhen(p *pipe) bool {
	return p.match(TK_WHEN)
}

func parseWhen(p *pipe) When {
	// pattern := WHEN BLOCK_OPEN guard {guard} BLOCK_CLOSE

	m := When{
		Key:   p.expect(`parseWhen`, TK_WHEN),
		Open:  p.expect(`parseWhen`, TK_BLOCK_OPEN),
		Cases: parseGuards(p),
	}

	if m.Cases == nil {
		err.Panic(
			errMsg("parseWhen", TK_GUARD_OPEN.String(), p.peek()),
			err.At(p.peek()),
		)
	}

	m.Close = p.expect(`parseWhen`, TK_BLOCK_CLOSE)
	return m
}
