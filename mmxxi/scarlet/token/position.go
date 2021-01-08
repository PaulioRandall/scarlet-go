package token

import (
	"fmt"
)

type (

	// Position represents a point in text but is decoupled from its exact
	// source, i.e. it does not hold a file or pointer to the source code.
	Position struct {
		Offset  int // Byte offset from start of text
		Line    int // Current line index
		ColByte int // Byte offset from start of the line
		ColRune int // Rune offset from start of line
	}

	// Snippet represents a range between two Positions within some text. When
	// operating with Snippets the source text must be the same for both,
	// however, they may overlap.
	Snippet struct {
		Start Position
		End   Position
	}

	// TextMarker provides functionality for moving progressing through some text.
	TextMarker Position
)

// String returns the position as a human readable string in the format:
//	offset[line:colByte/colRune]
func (p Position) String() string {
	return fmt.Sprintf("%d[%d:%d/%d]",
		p.Offset,
		p.Line,
		p.ColByte,
		p.ColRune,
	)
}

// From returns the position representing the beginning of the snippet.
func (s Snippet) From() Position {
	return s.Start
}

// To returns the position representing the end of the snippet.
func (s Snippet) To() Position {
	return s.End
}

// String returns a human readable string representation of the Snippet.
func (s Snippet) String() string {
	return fmt.Sprintf("%s -> %s", s.Start.String(), s.End.String())
}

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

// Position returns a position representing the end of 's' in source code
// assuming 's' starts at the TextMarker's current position.
func (tm *TextMarker) Position(s string) Position {
	p := TextMarker(tm.Snapshot())
	p.Advance(s)
	return Position(p)
}

// Snippet returns a Snippet representing 's' in source code assuming 's'
// starts at the TextMarker's current position.
func (tm *TextMarker) Snippet(s string) Snippet {
	start := tm.Snapshot()
	end := TextMarker(tm.Snapshot())
	end.Advance(s)
	return Snippet{
		Start: start,
		End:   Position(end),
	}
}

// Snapshot returns the current Position of the TextMarker.
func (tm *TextMarker) Snapshot() Position {
	return Position(*tm)
}
