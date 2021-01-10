package scanner

import (
	"io/ioutil"

	"github.com/PaulioRandall/scarlet-go/mmxxi/scarlet/scroll"
)

// ScrollReader reads of Scarlet source scrolls.
type ScrollReader interface {
	Line() int
	Slice(n int) string
	Read(n int) scroll.Snippet
	Peek(n int) scroll.Snippet
	More() bool
	At(i int) rune
	InRange(i int) bool
	Starts(text string) bool
	Contains(start int, text string) bool
}

type sReader struct {
	data []rune
	tm   scroll.TextMarker
}

// NewScrollReader returns an initialised scroll reader.
func NewScrollReader(s string) *sReader {
	return &sReader{
		data: []rune(s),
		tm:   scroll.TextMarker{},
	}
}

// Loads a scroll from a file.
func Load(filename string) (ScrollReader, error) {

	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}

	s := string(bytes)
	return NewScrollReader(s), nil
}

// Line returns the current line number.
func (sr *sReader) Line() int {
	return sr.tm.Line + 1
}

// Slice returns 'n' runes from the scroll as a string without advancing the
// text marker.
func (sr *sReader) Slice(n int) string {
	return string(sr.data[:n])
}

// Read reads 'n' runes from the scroll as a snippet.
func (sr *sReader) Read(n int) scroll.Snippet {
	text := string(sr.data[:n])
	sr.data = sr.data[n:]
	return sr.tm.Advance(text)
}

// Peek reads 'n' runes from the scroll as a snippet without progressing the
// text marker.
func (sr *sReader) Peek(n int) scroll.Snippet {
	text := string(sr.data[:n])
	return sr.tm.SliceSnippet(text)
}

// More returns true if there are runes yet to be read.
func (sr *sReader) More() bool {
	return len(sr.data) > 0
}

// At returns the rune at index 'i' from the scroll current position.
func (sr *sReader) At(i int) rune {
	return sr.data[i]
}

// InRange returns true if index 'i' is within the remaining ruens.
func (sr *sReader) InRange(i int) bool {
	return i < len(sr.data)
}

// Starts returns true if the remaining runes starts with 'text'.
func (sr *sReader) Starts(text string) bool {
	return sr.Contains(0, text)
}

// Contains returns true if the remaining runes contains the string 'text' at
// the index 'i'.
func (sr *sReader) Contains(start int, text string) bool {

	if start+len([]rune(text)) > len(sr.data) {
		return false
	}

	i := start
	for _, ru := range text {
		if sr.data[i] != ru {
			return false
		}
		i++
	}

	return true
}
