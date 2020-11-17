package position

type (

	// Position represents a point within a file.
	Position struct {
		filepath string
		offset   int // Byte offset from start of text
		line     int // Current line index
		col      int // Byte offset from start of the line
		colRune  int // Rune offset from start of the line
	}

	// Range represents a snippet between two points within a text source.
	Range struct {
		Position
		lineCount int // Line count
		len       int // Byte length from column start
		lenRune   int // Rune length from column start
	}

	// TextMarker represents a mutable Position within file text with
	// functionality for advancing through it.
	TextMarker Position
)

// Offset returns the byte offset within the file.
func (p Position) Offset() int {
	return p.offset
}

// Line returns the line index within the file.
func (p Position) Line() int {
	return p.line
}

// Col returns byte column index of the positions line within the file.
func (p Position) Col() int {
	return p.col
}

// ColRune returns the rune column index of the positions line within the file.
func (p Position) ColRune() int {
	return p.colRune
}

// LineCount returns the number of lines the range spans.
func (r Range) LineCount() int {
	return r.lineCount
}

// Len returns the byte length of the range.
func (r Range) Len() int {
	return r.len
}

// LenRune returns the rune length of the range.
func (r Range) LenRune() int {
	return r.lenRune
}

// Adv moves forward the number of bytes in 's'. For each linefeed '\n' in
// the string the line field is incremented and column values zeroed.
func (tm *TextMarker) Adv(s string) {
	tm.offset += len(s)
	for _, ru := range s {
		if ru == '\n' {
			tm.line++
			tm.col = 0
			tm.colRune = 0
		} else {
			tm.col += len(string(ru))
			tm.colRune++
		}
	}
}

// Pos returns the current Pos of the TextMarker.
func (tm *TextMarker) Pos() Position {
	return Position(*tm)
}

// Pos returns a new initialised Position.
func Pos(filepath string, offset, line, col, colRune int) Position {
	return Position{
		filepath: filepath,
		offset:   offset,
		line:     line,
		col:      col,
		colRune:  colRune,
	}
}

// Rng returns a new initialised Range.
func Rng(start Position, lineCount, len, lenRune int) Range {
	return Range{
		Position:  start,
		lineCount: lineCount,
		len:       len,
		lenRune:   lenRune,
	}
}
