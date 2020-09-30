package scanner

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token2/position"
)

type reader struct {
	position.TextMarker
	data   []rune
	remain int // Runes remaining
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

func (r *reader) read(runeCount int, newline bool) (position.Snippet, string) {

	val := r.slice(runeCount)
	snip := r.Snippet(val)

	r.data = r.data[runeCount:]
	r.remain = len(r.data)
	r.Advance(val, newline)

	return snip, val
}
