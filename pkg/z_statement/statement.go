package z_statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func Print(stats []Statement) {

	var s str

	s.append("[Statements]").
		newline()
	s.appendStats(1, stats)
	s.print()
}

type Statement interface {
	Token() Token
	String(indent int) string
}
