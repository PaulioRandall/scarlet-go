package z_matching

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

func ScanAll(s string) []Token {
	tks := readAllToks(s)
	return SanitiseAll(tks)
}

func readAllToks(s string) []Token {

	in := ReadAll(s)
	out := make([]Token, len(in))

	for i := range in {
		out[i] = in[i]
	}

	return out
}

type scanErr struct {
	msg  string
	line int
	col  int
	len  int
}

func err(s *symbols, colOffset int, msg string) error {
	return scanErr{
		line: s.line,
		col:  s.col + colOffset,
		msg:  msg,
	}
}

func (se scanErr) Error() string {
	return se.msg
}

func (se scanErr) Line() int {
	return se.line
}

func (se scanErr) Col() int {
	return se.col
}

func (se scanErr) Len() int {
	return se.len
}
