package formatter

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	//		"github.com/PaulioRandall/scarlet-go/eskarina/stages/a_scanner"
)

type line struct {
	original string
	head     *lexeme.Lexeme
	next     *line
}
