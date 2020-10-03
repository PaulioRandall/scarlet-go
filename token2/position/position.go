// Position package contains structs with receiving functions used to specify
// positions and snippets in source code without knowledge of the source codes
// location or format, i.e. structs do not store file paths, web addresses, or
// pointers to in-memory text strings.
package position

import (
	"fmt"
)

// Position represents a point in source code but is decoupled from its exact
// source, i.e. it does not hold a file or pointer to the source code.
type Position struct {
	Offset  int // Byte offset from start of text
	Line    int // Current line index
	ColByte int // Byte offset from start of the line
	ColRune int // Rune offset from start of line
}

// String returns the position as a human readable string in the format:
//	(offset[line:colByte/colRune])
func (p Position) String() string {
	return fmt.Sprintf("(%d[%d:%d/%d])",
		p.Offset,
		p.Line,
		p.ColByte,
		p.ColRune,
	)
}

// Snippet represents a range between two Positions within source code. The
// start Position is embedded for convenience.
type Snippet struct {
	Position // Start
	End      Position
}

// String returns a human readable string representation of the Snippet.
func (s Snippet) String() string {
	return fmt.Sprintf("%s -> %s", s.Position.String(), s.End.String())
}

// TextMarker represents a Position with functionality for advancing through
// source code.
type TextMarker Position

// Advance moves forward the number of bytes in 's'. For each linefeed '\n' in
// the string the line field is incremented and column values zeroed.
func (tm *TextMarker) Advance(s string) {
	tm.Offset += len(s)
	for _, ru := range s {
		if ru == '\n' {
			tm.Line++
			tm.ColByte = 0
			tm.ColRune = 0
		} else {
			tm.ColByte += len(string(ru))
			tm.ColRune++
		}
	}
}

// Snippet returns a Snippet representing 's' in source code assuming 's'
// starts at the TextMarker's current position.
func (tm *TextMarker) Snippet(s string) Snippet {
	start := tm.Snapshot()
	end := TextMarker(tm.Snapshot())
	end.Advance(s)
	return Snippet{
		Position: start,
		End:      Position(end),
	}
}

// Snapshot returns the current Position of the TextMarker.
func (tm *TextMarker) Snapshot() Position {
	return Position(*tm)
}
