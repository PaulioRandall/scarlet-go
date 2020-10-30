package parser

import (
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type context struct {
	*token.LexItr
	parent *context
}
