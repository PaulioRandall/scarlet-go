package parser2

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type context struct {
	*token.LexItr
	parent *context
}
