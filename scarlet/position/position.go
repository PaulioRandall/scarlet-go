package position

import (
	"fmt"
)

type (
	// Position represents a point within a text source.
	Position interface {
		Offset() int  // Byte offset from start of text
		Line() int    // Current line index
		Col() int     // Byte offset from start of the line
		ColRune() int // Rune offset from start of the line
	}

	// Range represents a snippet between two points within a text source.
	Range interface {
		From() Position
		To() Position
	}

	// Pos is an immutable implementation of Position.
	Pos struct {
		offset  int
		line    int
		col     int
		colRune int
	}

	// TextMarker represents a mutable Position within some text with
	// functionality for progressing through the text.
	TextMarker struct {
		pos Pos
	}
)

// Pos implements the Position interface

func (p Pos) Offset() int {
	return p.offset
}

func (p Pos) Line() int {
	return p.line
}

func (p Pos) Col() int {
	return p.col
}

func (p Pos) ColRune() int {
	return p.colRune
}

// Adv moves forward the number of bytes in 's'. For each linefeed '\n' in
// the string the line field is incremented and column values zeroed.
func (tm *TextMarker) Adv(s string) {
	tm.pos.offset += len(s)
	for _, ru := range s {
		if ru == '\n' {
			tm.pos.line++
			tm.pos.col = 0
			tm.pos.colRune = 0
		} else {
			tm.pos.col += len(string(ru))
			tm.pos.colRune++
		}
	}
}

// Copy returns a copy of the TextMarker.
func (tm *TextMarker) Copy() TextMarker {
	return *tm
}

// Pos returns the current Pos of the TextMarker.
func (tm *TextMarker) Pos() Pos {
	return tm.pos
}

// Make returns a new initialised Pos.
func Make(offset, line, col, colRune int) Pos {
	return Pos{
		offset:  offset,
		line:    line,
		col:     col,
		colRune: colRune,
	}
}

// PosStr returns human readable string of a Position in the format:
// offset(line:colByte/colRune)
func PosStr(p Position) string {
	return fmt.Sprintf("%d(%d:%d/%d)",
		p.Offset(),
		p.Line(),
		p.Col(),
		p.ColRune(),
	)
}

// RngStr returns human readable string of a Range in the format:
// [offset(line:colByte/colRune):offset(line:colByte/colRune)]
func RngStr(r Range) string {
	return fmt.Sprintf("[%s -> %s]", r.From(), r.To())
}
