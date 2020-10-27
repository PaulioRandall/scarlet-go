package cmd

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

// More returns true if there are more arguments to the right of the item index.
func (a *Args) More() bool {
	return a.Idx < a.Size
}

// Shift increments the argument pointer one item to the right then returns the
// item at the new index.
func (a *Args) Shift() string {
	v := a.Peek()
	a.Idx++
	return v
}

// ShiftDefault increments the argument pointer one item to the right then
// returns the item at the new index. If the there is no items remaining the
// the default value is returned.
func (a *Args) ShiftDefault(def string) string {
	if !a.More() {
		return def
	}
	v := a.Vals[a.Idx]
	a.Idx++
	return v
}

// Accept increments the index pointer and returns true if 's' exactly matches
// the next item.
func (a *Args) Accept(s string) bool {
	if a.More() && strings.ToLower(a.Peek()) == strings.ToLower(s) {
		a.Shift()
		return true
	}
	return false
}

// Peek returns the next argument without incrementing the index.
func (a *Args) Peek() string {
	if !a.More() {
		panic("Out of range") // TODO: Create better error message
	}
	return a.Vals[a.Idx]
}
