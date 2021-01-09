package scroll

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
		Text  string
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
	return fmt.Sprintf("%q %s -> %s",
		s.Text,
		s.Start.String(),
		s.End.String(),
	)
}

// Pos returns the current Position of the TextMarker.
func (tm TextMarker) Pos() Position {
	return Position(tm)
}

// EndOf returns a Position representing the end of 's' assuming 's' starts at
// the TextMarker's current position. No advancing takes place.
func (tm TextMarker) EndOf(s string) Position {
	end := TextMarker(tm.Pos())
	end.Advance(s)
	return Position(end)
}

// Advance increments the number of bytes in 's'. For each linefeed '\n'
// the line field is incremented and column values zeroed. A snippet
// representing 's' is returned.
func (tm *TextMarker) Advance(s string) Snippet {
	start := tm.Pos()

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

	return Snippet{
		Start: start,
		End:   tm.Pos(),
		Text:  s,
	}
}

// SliceSnippet returns the snippet of 's'.
func (tm *TextMarker) SliceSnippet(s string) Snippet {
	cp := TextMarker(tm.Pos())
	return cp.Advance(s)
}
