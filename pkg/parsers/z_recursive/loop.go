package z_recursive

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func isLoop(p *pipe) bool {
	return p.match(LOOP)
}

func parseLoop(p *pipe) Loop {
	// pattern := LOOP ID guard

	return Loop{
		Open:     p.expect(`parseLoop`, LOOP),
		IndexVar: p.expect(`parseLoop`, IDENTIFIER),
		Guard:    parseGuard(p),
	}
}
