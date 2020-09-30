package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/token"
)

type reader struct {
	data    []rune
	remain  int // Runes remaining
	offset  int // Byte offset from start of source file
	line    int // Current line index
	colByte int // Byte offset from start of the line
	colRune int // Rune offset from start of line
}

func newReader(s string) *reader {
	r := &reader{}
	r.data = []rune(s)
	r.remain = len([]rune(r.data))
	return r
}

func (r *reader) more() bool {
	return r.remain > 0
}

func (r *reader) at(i int) rune {
	return r.data[i]
}

func (r *reader) inRange(i int) bool {
	return i < r.remain
}

func (r *reader) starts(s string) bool {
	return r.contains(0, s)
}

func (r *reader) contains(start int, s string) bool {

	dataSize := len(r.data)
	if start > dataSize {
		e := fmt.Errorf(
			"Start index i out of range, given %d, want <%d", start, dataSize)
		panic(e)
	}

	if start+len([]rune(s)) > dataSize {
		return false
	}

	i := start
	for _, ru := range s {
		if r.data[i] != ru {
			return false
		}
		i++
	}

	return true
}

func (r *reader) slice(runeCount int) string {
	return string(r.data[:runeCount])
}

func (r *reader) read(runeCount int, newline bool) (token.Snippet, string) {

	val := r.slice(runeCount)
	byteCount := len(val)
	snip := r.snippet(byteCount, runeCount)

	r.offset += byteCount
	r.data = r.data[runeCount:]
	r.remain = len(r.data)

	if newline {
		r.line++
		r.colByte = 0
		r.colRune = 0
	} else {
		r.colByte += byteCount
		r.colRune += runeCount
	}

	return snip, val
}

func (r *reader) snippet(byteCount, runeCount int) token.Snippet {
	return token.Snippet{
		Position: token.Position{
			SrcOffset: r.offset,
			LineIdx:   r.line,
			ColByte:   r.colByte,
			ColRune:   r.colRune,
		},
		End: token.Position{
			SrcOffset: r.offset + byteCount,
			LineIdx:   r.line,
			ColByte:   r.colByte + byteCount,
			ColRune:   r.colRune + runeCount,
		},
	}
}
