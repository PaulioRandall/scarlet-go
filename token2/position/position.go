package position

import (
	"fmt"
)

type Position struct {
	Offset  int // Bytes
	Line    int // Index
	ColByte int // Index
	ColRune int // Index
}

func (p Position) String() string {
	return fmt.Sprintf("[%d]%d/%d/%d",
		p.Line,
		p.Offset,
		p.ColByte,
		p.ColRune,
	)
}

type Snippet struct {
	Position
	End Position
}

func (s Snippet) String() string {
	return fmt.Sprintf("%s -> %s", s.Position.String(), s.End.String())
}

type TextMarker struct {
	Offset  int // Byte offset from start of text
	Line    int // Current line index
	ColByte int // Byte offset from start of the line
	ColRune int // Rune offset from start of line
}

func (tm *TextMarker) Advance(v string, newline bool) {

	byteCount := len(v)
	runeCount := len([]rune(v))

	tm.Offset += byteCount
	if newline {
		tm.Line++
		tm.ColByte = 0
		tm.ColRune = 0
	} else {
		tm.ColByte += byteCount
		tm.ColRune += runeCount
	}
}

func (tm *TextMarker) Snippet(v string) Snippet {

	byteCount := len(v)
	runeCount := len([]rune(v))

	return Snippet{
		Position: tm.Position(),
		End: Position{
			Offset:  tm.Offset + byteCount,
			Line:    tm.Line,
			ColByte: tm.ColByte + byteCount,
			ColRune: tm.ColRune + runeCount,
		},
	}
}

func (tm *TextMarker) Position() Position {
	return Position{
		Offset:  tm.Offset,
		Line:    tm.Line,
		ColByte: tm.ColByte,
		ColRune: tm.ColRune,
	}
}
