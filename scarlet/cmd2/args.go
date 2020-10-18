package cmd2

import (
	"strings"
)

// Args represents a scroll's arguments.
type Args struct {
	Vals []string
	Size int
	Idx  int
}

// NewArgs returns a new arguement iterator.
func NewArgs(vals []string) *Args {
	return &Args{
		Vals: vals,
		Size: len(vals),
	}
}

// Empty returns true if there are no more arguments to the right of the item
// index.
func (a *Args) Empty() bool {
	return a.Idx+1 >= a.Size
}

// More returns true if there are more arguments to the right of the item index.
func (a *Args) More() bool {
	return a.Idx+1 < a.Size
}

// Shift increments the argument pointer one item to the right then returns the
// item at the new index.
func (a *Args) Shift() string {
	if !a.inRange(1) {
		panic("Out of range") // TODO: Create better error message
	}
	a.Idx++
	return a.Vals[a.Idx]
}

// Unshift decrements the argument pointer one item to the left then returns the
// item at the new index.
func (a *Args) Unshift() string {
	if !a.inRange(-1) {
		panic("Out of range") // TODO: Create better error message
	}
	a.Idx--
	return a.Vals[a.Idx]
}

// ShiftDefault increments the argument pointer one item to the right then
// returns the item at the new index. If the there is no items remaining the
// the default value is returned.
func (a *Args) ShiftDefault(def string) string {
	if a.inRange(1) {
		return def
	}
	a.Idx++
	return a.Vals[a.Idx]
}

// Accept increments the index pointer and returns true if 's' exactly matches
// the next item.
func (a *Args) Accept(s string) bool {
	if a.inRange(1) && a.Peek() == s {
		a.Shift()
		return true
	}
	return false
}

// Peek returns the next argument without incrementing the index.
func (a *Args) Peek() string {
	if !a.inRange(1) {
		panic("Out of range") // TODO: Create better error message
	}
	return a.Vals[a.Idx]
}

// IsOption returns true if the argument is an option.
func (args *Args) isOption() bool {
	return strings.HasPrefix(args.Peek(), "-")
}

func (a *Args) inRange(offset int) bool {
	return a.Idx+offset >= 0 && a.Idx+offset < a.Size
}
