package position

type (

	// Position represents a point within a file.
	Position struct {
		path    string // Filepath to the source text file
		offset  int    // Byte offset from start of text
		line    int    // Current line index
		byteCol int    // Byte offset from start of the line
		runeCol int    // Rune offset from start of the line
	}

	// Range represents a snippet between two points within a text source.
	Range struct {
		Position
		lineCount int // Line count
		byteLen   int // Byte length
	}

	// TextMarker represents a mutable Position within a text file with
	// functionality for advancing through it.
	TextMarker Position
)

// Path returns the filepath to the file.
func (p Position) Path() string {
	return p.path
}

// Offset returns the byte offset within the file.
func (p Position) Offset() int {
	return p.offset
}

// Line returns the line index within the file.
func (p Position) Line() int {
	return p.line
}

// ByteCol returns byte column index of the positions line within the file.
func (p Position) ByteCol() int {
	return p.byteCol
}

// RuneCol returns the rune column index of the positions line within the file.
func (p Position) RuneCol() int {
	return p.runeCol
}

// LineCount returns the number of lines the range spans.
func (r Range) LineCount() int {
	return r.lineCount
}

// ByteLen returns the byte length of the range.
func (r Range) ByteLen() int {
	return r.byteLen
}

// Adv moves forward the number of bytes in 's'. For each linefeed '\n' in
// the string the line field is incremented and column values zeroed.
func (tm *TextMarker) Adv(s string) {
	tm.offset += len(s)
	for _, ru := range s {
		if ru == '\n' {
			tm.line++
			tm.byteCol = 0
			tm.runeCol = 0
		} else {
			tm.byteCol += len(string(ru))
			tm.runeCol++
		}
	}
}

// RangeOf returns a Range between the current position and the end of 's'.
func (tm *TextMarker) RangeOf(s string) Range {
	start, end := *tm, *tm
	end.Adv(s)
	return Range{
		Position:  Position(start),
		lineCount: end.line - start.line + 1,
		byteLen:   end.offset - start.offset,
	}
}

// PosAfter returns the position after advancing 's'.
func (tm *TextMarker) PosAfter(s string) Position {
	p := *tm
	p.Adv(s)
	return Position(p)
}

// Pos returns the current Pos of the TextMarker.
func (tm *TextMarker) Pos() Position {
	return Position(*tm)
}

// Pos returns a new initialised Position.
func Pos(path string, offset, line, byteCol, runeCol int) Position {
	return Position{
		path:    path,
		offset:  offset,
		line:    line,
		byteCol: byteCol,
		runeCol: runeCol,
	}
}

// Rng returns a new initialised Range.
func Rng(start Position, lineCount, byteLen int) Range {
	return Range{
		Position:  start,
		lineCount: lineCount,
		byteLen:   byteLen,
	}
}
