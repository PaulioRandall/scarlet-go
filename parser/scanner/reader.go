package scanner

import (
	"fmt"
)

type reader struct {
	data []rune
	size int
	line int // Track current line
	col  int // Track current column within line
}

func (r *reader) more() bool {
	return r.size > 0
}

func (r *reader) at(i int) rune {
	return r.data[i]
}

func (r *reader) inRange(i int) bool {
	return i < r.size
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

func (r *reader) slice(size int) string {
	return string(r.data[:size])
}

func (r *reader) read(size int, newline bool) (line, col int, s string) {

	line, col = r.line, r.col
	if newline {
		r.line++
		r.col = 0
	} else {
		r.col += size
	}

	s = string(r.data[:size])
	r.data = r.data[size:]
	r.size = len(r.data)
	return
}
