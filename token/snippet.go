package token

import (
	"fmt"
)

// Snippet represents the location of a snippet of text
// (UTF-8 characters/runes) within a file.
type Snippet interface {
	Line() int  // Index of the line in the file
	Start() int // Column index where the snippet begins
	End() int   // Column index after the last character in the snippet
	String() string
}

// snip is an implementation of Snippet.
type snip struct {
	line  int
	start int
	end   int
}

// NewSnippet creates a new Snippet.
func NewSnippet(line, start, end int) Snippet {
	return snip{
		line:  line,
		start: start,
		end:   end,
	}
}

// Line satisfies the Snippet interface.
func (s snip) Line() int {
	return s.line
}

// Start satisfies the Snippet interface.
func (s snip) Start() int {
	return s.start
}

// End satisfies the Snippet interface.
func (s snip) End() int {
	return s.end
}

// String returns a string representation of the entity.
func (s snip) String() string {
	return fmt.Sprintf(`line %d (%d:%d)`, s.line, s.start, s.end)
}
