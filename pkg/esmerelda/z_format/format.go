package format

import (
	"strings"
)

type line struct {
	idx  int
	raw  string
	prev *line
	next *line
}

func Format(s string) string {

	lineEndings := detectLineEndings(s)
	lines := splitLines(s)

	// https://en.wikipedia.org/wiki/Parsing_expression_grammar

	// Use a functional approach:
	// 1: Correct line endings so all equal to 'lineEndings'
	// 2: Remove all redundant whitespace
	// 3: Remove multiple empty lines
	// 4: Insert single space after value separators if not a newline, i.e. ','
	// 5: Remove empty lines between list items
	// 6: Indent for multiline statements (except initiating line and final ')')
	// 7: Align comments for consecutive lines with comments

	_, _ = lineEndings, lines
	return s
}

func splitLines(s string) *line {

	var first line
	curr := &first

	for i, l := range strings.Split(s, "\n") {

		if i != 0 {
			curr.next = &line{
				prev: curr,
			}
			curr = curr.next
		}

		curr.idx = i
		curr.raw = l
	}

	return &first
}

func detectLineEndings(s string) string {

	var found bool
	size := len(s)
	i := 0

	for ; i < size; i++ {
		if s[i] == '\r' || s[i] == '\n' {
			found = true
			break
		}
	}

	if !found {
		return ""
	}

	if s[i] == '\n' {
		return "\n"
	}

	if i+1 >= size || s[i+1] != '\n' {
		panic("Expected linefeed '\n'")
	}

	return "\r\n"
}
