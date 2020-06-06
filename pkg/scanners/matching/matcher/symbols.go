package matcher

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/pkg/err"
)

type RuneMatcher func(index int, value rune) (match bool, e error)

type Symbols struct {
	runes  []rune
	line   int
	col    int
	offset int
}

func (s *Symbols) Pos() (int, int) {
	return s.line, s.col
}

func (s *Symbols) Empty() bool {
	return s.Remaining() <= 0
}

func (s *Symbols) Remaining() int {
	return len(s.runes) - s.offset
}

func (s *Symbols) Has(n int) bool {
	return (s.Remaining() - n) > -1
}

func (s *Symbols) At(i int) (rune, error) {

	var e error

	i, e = s.offsetIndex(i, false)
	if e != nil {
		return rune(0), e
	}

	return s.runes[i], nil
}

func (s *Symbols) Slice(start, end int) (string, error) {

	var e error

	start, e = s.offsetIndex(start, false)
	if e != nil {
		return ``, e
	}

	end, e = s.offsetIndex(end, true)
	if e != nil {
		return ``, e
	}

	return string(s.runes[start:end]), nil
}

func (s *Symbols) Match(start int, str string) (bool, error) {

	var e error

	start, e = s.offsetIndex(start, false)
	if e != nil {
		return false, e
	}

	haystack := s.runes[start:]

	if len(str) > len(haystack) {
		return false, nil
	}

	for i, ru := range str {
		if haystack[i] != ru {
			return false, nil
		}
	}

	return true, nil
}

func (s *Symbols) CountWhile(start int, f RuneMatcher) (int, error) {

	var e error

	start, e = s.offsetIndex(start, true)
	if e != nil {
		return 0, e
	}

	size := s.Remaining()
	for i := start; i < size; i++ {

		match, e := f(i, s.runes[i])
		if e != nil {
			return 0, e
		}

		if !match {
			return i - start, nil
		}
	}

	return size - start, nil
}

func (s *Symbols) IsNewline(index int) (bool, int) {
	count := s.countNewlineTerminals(index)
	return count > 0, count
}

func (s *Symbols) offsetIndex(index int, includeLen bool) (int, error) {

	i := index + s.offset
	rem := s.Remaining()

	if i < 0 {
		goto ERROR
	}

	if i > rem {
		goto ERROR
	}

	if !includeLen && i == rem {
		goto ERROR
	}

	return i, nil

ERROR:
	return 0, err.New(
		fmt.Sprintf(
			"Index out of range, given %d, but got [%d:%d]",
			index, 0, s.Remaining(),
		),
		err.Pos(s.line, s.col),
	)
}

func (s *Symbols) read(runeCount int) (string, error) {

	r, e := s.Slice(0, runeCount)
	if e != nil {
		return ``, e
	}

	for i := 0; i < runeCount; i++ {

		switch s.countNewlineTerminals(i) {
		case 2:
			i++
			fallthrough

		case 1:
			s.line++
			s.col = 0

		case 0:
			s.col++
		}
	}

	s.offset += len(r)
	return r, nil
}

func (s *Symbols) countNewlineTerminals(index int) int {

	if !s.Has(1) {
		return 0
	}

	match, _ := s.Match(index, "\n")
	if match {
		return 1
	}

	if !s.Has(2) {
		return 0
	}

	match, _ = s.Match(index, "\r\n")
	if match {
		return 2
	}

	return 0
}
