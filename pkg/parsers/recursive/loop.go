package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func isLoop(p *pipe) bool {
	return p.match(token.LOOP)
}

func parseLoop(p *pipe) st.Loop {
	// pattern := LOOP ID guard

	return st.Loop{
		Open:     p.expect(`parseLoop`, token.LOOP),
		IndexVar: p.expect(`parseLoop`, token.ID),
		Guard:    parseGuard(p),
	}
}
