package position

type (

	// Range represents some text between two points within a text file.
	Range struct {
		Rpath   string // Filepath to the source text file
		Roffset int    // Byte offset from start of text
		Rline   int    // Current line index
		Rcol    int    // Byte offset from start of the line
		Rlen    int    // Byte length
	}

	// TextMarker represents a mutable Range within a text file with
	// functionality for advancing through it. 'Range.Rlen' is ignored.
	TextMarker Range
)

// Path returns the filepath to the file.
func (r Range) Path() string {
	return r.Rpath
}

// Offset returns the byte offset within the file.
func (r Range) Offset() int {
	return r.Roffset
}

// Line returns the line index within the file.
func (r Range) Line() int {
	return r.Rline
}

// Col returns column index of the positions line within the file.
func (r Range) Col() int {
	return r.Rcol
}

// Len returns the byte length of the range.
func (r Range) Len() int {
	return r.Rlen
}

// Adv moves forward the number of bytes in 's'. For each linefeed '\n' in
// the string the line field is incremented and column values zeroed.
func (tm *TextMarker) Adv(s string) {
	for _, b := range []byte(s) {
		tm.Roffset++
		if b == byte('\n') {
			tm.Rline++
			tm.Rcol = 0
		} else {
			tm.Rcol++
		}
	}
}

// Range returns the current position of the TextMarker.
func (tm *TextMarker) Range() Range {
	return Range(*tm)
}

// RangeOf returns a Range between the current position and the end of 's'.
func (tm *TextMarker) RangeOf(s string) Range {
	rng := *tm
	rng.Rlen = len(s)
	return Range(rng)
}
