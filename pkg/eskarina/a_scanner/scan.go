package scanner

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
)

func ScanAll(s string) *lexeme.Lexeme {
	return nil
}

func scanNext(rr *runeReader) {

	switch rr.peek() {
	case '\r':
	case '\n':
	}
}

func scanNewline(rr *runeReader) *lexeme.Lexeme {

	if rr.accept('\n') {
		return nil
	}

	return nil
}
