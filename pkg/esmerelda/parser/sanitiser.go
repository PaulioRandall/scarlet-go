package parser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/token"
)

type sanitiser struct {
	stream TokenStream
	buff   token.Token
	back   token.Token
	prev   token.Token
}

func newSanitiser(stream TokenStream) sanitiser {
	s := sanitiser{stream, nil, nil, nil}
	s.bufferNext()
	return s
}

func (s *sanitiser) bufferNext() {

	s.prev = s.buff

	if s.back != nil {
		s.buff = s.back
		s.back = nil
		return
	}

	buff := s.stream.Next()
	for buff != nil && s._ignore(buff, s.prev) {
		buff = s.stream.Next()
	}

	if buff == nil {
		s.buff = nil
		return
	}

	s.buff = s._format(buff)
}

func (s *sanitiser) empty() bool {
	return s.buff == nil
}

func (s *sanitiser) peek() token.Token {
	return s.buff
}

func (s *sanitiser) next() token.Token {
	s.bufferNext()
	return s.prev
}

func (s *sanitiser) past() token.Token {
	return s.prev
}

func (s *sanitiser) backup() {

	if s.back != nil {
		panic("PROGRAMMERS ERROR! Cannot backtrack twice in a row")
	}

	if s.prev == nil {
		panic("PROGRAMMERS ERROR! Cannot backtrack past the start of the stream")
	}

	s.back = s.buff
	s.buff = s.prev
	s.prev = nil
}

func (s *sanitiser) _ignore(next, prev token.Token) bool {

	ty := next.Type()

	switch {
	case ty == token.TK_COMMENT:
		return true

	case ty == token.TK_WHITESPACE:
		return true

	case ty != token.TK_TERMINATOR:
		return false

		// next must be a TERMINATOR
	case prev == nil: // Ignore TERMINATORs at start of script
		return true

	case prev.Type() == token.TK_DELIMITER: // Allow "NEWLINE" after delimiter
		return true

	case prev.Type() == token.TK_BLOCK_OPEN: // Allow "NEWLINE" after block start
		return true

	case prev.Type() == token.TK_PAREN_OPEN: // Allow "NEWLINE" after paren start
		return true

	case prev.Type() == token.TK_TERMINATOR: // Ignore successive TERMINATORs
		return true
	}

	return false
}

func (s *sanitiser) _format(tk token.Token) token.Token {

	if tk.Type() == token.TK_NEWLINE {
		line, col := tk.Begin()
		return token.NewToken(token.TK_TERMINATOR, tk.Value(), line, col)
	}

	return tk
}
