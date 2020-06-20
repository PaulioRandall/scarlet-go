package statement

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/gytha/token"
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
