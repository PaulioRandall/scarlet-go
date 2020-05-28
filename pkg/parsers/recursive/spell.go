package recursive

import (
	//	"github.com/PaulioRandall/scarlet-go/pkg/err"
	//. "github.com/PaulioRandall/scarlet-go/pkg/statement"
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

func isSpell(p *pipe) bool {
	return p.match(SPELL)
}
