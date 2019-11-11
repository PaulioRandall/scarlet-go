package where

import (
	"fmt"
)

// Where represents a snippet of text (UTF-8 characters/runes), within a file.
type Where interface {
	Line() int  // Index of the line in the file
	Start() int // Column index where the snippet begins
	End() int   // Column index after the last character in the snippet
	String() string
}

// whe is an implementation of Where.
type whe struct {
	line  int
	start int
	end   int
}

// Line satisfies the Where interface.
func (w whe) Line() int {
	return w.line
}

// Start satisfies the Where interface.
func (w whe) Start() int {
	return w.start
}

// End satisfies the Where interface.
func (w whe) End() int {
	return w.end
}

// String returns a string representation of the entity.
func (w whe) String() string {
	return fmt.Sprintf(`line %d (%d:%d)`, w.line, w.start, w.end)
}

// New makes a structure that implements the Where interface.
func New(line, start, end int) Where {
	return whe{
		line:  line,
		start: start,
		end:   end,
	}
}
