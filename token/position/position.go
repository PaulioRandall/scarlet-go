// Position package contains structs with receiving functions used to specify
// positions and snippets in source code without knowledge of the source codes
// location or format, i.e. structs do not store file paths, web addresses, or
// pointers to in-memory text strings.
package position

import (
	"fmt"
)

type (

	// UTF8Pos represents a point in text but is decoupled from its exact
	// source, i.e. it does not hold a file or pointer to the source code.
	UTF8Pos struct {
		Offset  int // Byte offset from start of text
		Line    int // Current line index
		ColByte int // Byte offset from start of the line
		ColRune int // Rune offset from start of line
	}

	// Snippet represents a range between two Positions within some text. The
	// start UTF8Pos is embedded for convenience. When operating with Snippets
	// the source text must be the same for both, however, they may overlap.
	Snippet struct {
		UTF8Pos // Start
		End     UTF8Pos
	}

	// TextMarker represents a moving UTF8Pos within some text with
	// functionality for progressing through the text.
	TextMarker UTF8Pos
)

// First returns the UTF8Pos that is nearer the start of the source text. The
// positions provided must be valid and from the same source text for the
// result to be meaningful.
func First(a, b UTF8Pos) UTF8Pos {
	if a.Offset < b.Offset {
		return a
	}
	return b
}

// Last returns the UTF8Pos that is nearer the end of the source text. The
// positions provided must be valid and from the same source text for the
// result to be meaningful.
func Last(a, b UTF8Pos) UTF8Pos {
	if a.Offset > b.Offset {
		return a
	}
	return b
}

// String returns the position as a human readable string in the format:
//	offset[line:colByte/colRune]
func (p UTF8Pos) String() string {
	return fmt.Sprintf("%d[%d:%d/%d]",
		p.Offset,
		p.Line,
		p.ColByte,
		p.ColRune,
	)
}

// SuperSnippet returns the smallest Snippet that contains the Snippets in
// 'a' and 'b', i.e. a Snippet with the nearest and furthest most UTF8Pos from
// the beginning of the source text. The Snippets provided must be valid and
// from the same source text for the result to be meaningful.
func SuperSnippet(a, b Snippet) Snippet {
	return Snippet{
		UTF8Pos: First(a.UTF8Pos, b.UTF8Pos),
		End:     Last(a.End, b.End),
	}
}

// From returns the UTF8Pos representing the beginning of the snippet.
func (s Snippet) From() UTF8Pos {
	return s.UTF8Pos
}

// To returns the UTF8Pos representing the end of the snippet.
func (s Snippet) To() UTF8Pos {
	return s.End
}

// String returns a human readable string representation of the Snippet.
func (s Snippet) String() string {
	return fmt.Sprintf("%s -> %s", s.UTF8Pos.String(), s.End.String())
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

// Position returns a UTF8Pos representing the end of 's' in source code
// assuming 's' starts at the TextMarker's current position.
func (tm *TextMarker) Position(s string) UTF8Pos {
	pos := TextMarker(tm.Snapshot())
	pos.Advance(s)
	return UTF8Pos(pos)
}

// Snippet returns a Snippet representing 's' in source code assuming 's'
// starts at the TextMarker's current position.
func (tm *TextMarker) Snippet(s string) Snippet {
	start := tm.Snapshot()
	end := TextMarker(tm.Snapshot())
	end.Advance(s)
	return Snippet{
		UTF8Pos: start,
		End:     UTF8Pos(end),
	}
}

// Snapshot returns the current UTF8Pos of the TextMarker.
func (tm *TextMarker) Snapshot() UTF8Pos {
	return UTF8Pos(*tm)
}
