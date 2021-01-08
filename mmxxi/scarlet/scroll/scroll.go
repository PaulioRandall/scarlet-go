package scroll

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Scroll represents a Scarlet scroll or source file.
type Scroll struct {
	Name string // Name of the file without extension
	Dir  string // Directory containing the file
	Path string // Path to the file including name and extension
}

// Make returns a new scroll from a file path.
func Make(path string) (s Scroll, e error) {

	s.Path, e = filepath.Abs(path)
	if e != nil {
		return s, e
	}

	s.Name = filepath.Base(path)
	s.Name = strings.TrimSuffix(s.Name, filepath.Ext(s.Name))
	s.Dir = filepath.Dir(path)

	return s, nil
}

// Reader returns a new scroll reader.
func (s Scroll) Reader() ScrollReader {
	return MakeReader(s.Path)
}

// ScrollReader represents a reader of Scarlet source.
type ScrollReader struct {
	data []rune
	tm   *TextMarker
}

// Make returns an initialised scroll.
func MakeReader(s string) ScrollReader {
	return ScrollReader{
		data: []rune(s),
		tm:   &TextMarker{},
	}
}

// Loads a scroll from a file.
func Load(filename string) (ScrollReader, error) {

	bytes, e := ioutil.ReadFile(filename)
	if e != nil {
		return ScrollReader{}, e
	}

	s := string(bytes)
	return MakeReader(s), nil
}

// Read reads 'n' runes from the scroll as a snippet.
func (sr ScrollReader) Read(n int) Snippet {
	text := string(sr.data[:n])
	sr.data = sr.data[n:]
	return sr.tm.Advance(text)
}

// Peek reads 'n' runes from the scroll as a snippet without progressing the
// text marker.
func (sr ScrollReader) Peek(n int) Snippet {
	text := string(sr.data[:n])
	return sr.tm.SliceSnippet(text)
}

// More returns true if there are runes yet to be read.
func (sr ScrollReader) More() bool {
	return len(sr.data) > 0
}

// At returns the rune at index 'i' from the scroll current position.
func (sr ScrollReader) At(i int) rune {
	return sr.data[i]
}

// InRange returns true if index 'i' is within the remaining ruens.
func (sr ScrollReader) InRange(i int) bool {
	return i < len(sr.data)
}

// Starts returns true if the remaining runes starts with 'text'.
func (sr ScrollReader) Starts(text string) bool {
	return sr.Contains(0, text)
}

// Contains returns true if the remaining runes contains the string 'text' at
// the index 'i'.
func (sr ScrollReader) Contains(start int, text string) bool {

	dataSize := len(sr.data)
	if start >= dataSize {
		e := fmt.Errorf(
			"Start index out of range, given %d, want <%d", start, dataSize)
		panic(e)
	}

	if start+len([]rune(text)) > dataSize {
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
