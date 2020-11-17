package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/scarlet/position"
	"github.com/PaulioRandall/scarlet-go/scarlet/token"
)

type (

	// Position represents a point within a text source.
	Position interface {
		Offset() int  // Byte offset from start of text
		Line() int    // Current line index
		Col() int     // Byte offset from start of the line
		ColRune() int // Rune offset from start of the line
	}

	scanErr struct {
		msg string
		Position
	}
)

// scanErr implements the error interface.

func (e scanErr) Error() string {
	return e.msg
}

func err(p token.UTF8Pos, msg string, args ...interface{}) scanErr {
	return scanErr{
		// TODO: Filename
		Position: position.Pos("", p.Offset, p.Line, p.ColByte, p.ColRune),
		msg:      fmt.Sprintf(msg, args...),
	}
}
