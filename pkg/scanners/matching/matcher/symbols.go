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

func (s *Symbols) At(i int) rune {

	if i < 0 || i >= s.Remaining() {
		return rune(0)
	}

	return s.runes[i+s.offset]
}

func (s *Symbols) Slice(start, end int) (string, error) {

	runes := s.getRemaining()

	e := s.checkStartIndex(start)
	if e != nil {
		return ``, e
	}

	e = s.checkEndIndex(end)
	if e != nil {
		return ``, e
	}

	return string(runes[start:end]), nil
}

func (s *Symbols) Match(start int, str string) (bool, error) {

	e := s.checkStartIndex(start)
	if e != nil {
		return false, e
	}

	haystack := s.getRemaining()[start:]

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

	e := s.checkStartIndex(start)
	if e != nil {
		return 0, e
	}

	runes := s.getRemaining()
	size := len(runes)

	if start >= size {
		return 0, nil
	}

	for i := start; i < size; i++ {

		match, e := f(i, runes[i])
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

func (s *Symbols) getRemaining() []rune {
	return s.runes[s.offset:]
}

func (s *Symbols) newError(msg string, args ...interface{}) error {
	return err.New(
		fmt.Sprintf(msg, args...),
		err.Pos(s.line, s.col),
	)
}

func (s *Symbols) checkStartIndex(start int) error {

	if start < 0 {
		return s.newError("Start should be 0 or more, not %d", start)
	}

	size := s.Remaining()
	if start >= size {
		return s.newError("Start should be %d or less, not %d", size, start)
	}

	return nil
}

func (s *Symbols) checkEndIndex(end int) error {

	if end < 0 {
		return s.newError("End should be 0 or more, not %d", end)
	}

	size := s.Remaining()
	if end > size {
		return s.newError("End should be %d or less, not %d", size, end)
	}

	return nil
}
