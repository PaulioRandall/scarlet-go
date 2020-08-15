package fmtr

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/PaulioRandall/scarlet-go/parser/a_scanner"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

func formatFile(filename string) error {

	c, e := readFile(filename)
	if e != nil {
		return e
	}

	format(c.Iterator())
	return writeFile(filename, c.Iterator())
}

func readFile(filename string) (*lexeme.Container, error) {

	b, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}

	return scanner.ScanStr(string(b))
}

func writeFile(filename string, itr *lexeme.Iterator) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()

	for itr.Next() {
		_, e = f.WriteString(itr.Curr().Raw)
		if e != nil {
			return e
		}
	}

	return nil
}

func format(itr *lexeme.Iterator) {

	trimWhiteSpace(itr)
	itr.Restart()

	stripUselessLines(itr)
	itr.Restart()

	insertSeparatorSpaces(itr)
	itr.Restart()

	insertAssignmentSpaces(itr)
	itr.Restart()

	insertOperatorSpaces(itr)
	itr.Restart()

	insertCommentSpaces(itr)
	itr.Restart()

	unifyLineEndings(itr)
	itr.Restart()

	indentLines(itr)
	itr.Restart()

	updatePositions(itr)
	itr.Restart()

	alignComments(itr)
	itr.Restart()
}

func trimWhiteSpace(itr *lexeme.Iterator) {

	whitespace := func(v lexeme.View) bool {
		return v.Curr().Tok == lexeme.SPACE
	}

	for itr.JumpToNext(whitespace) {
		itr.Remove()
	}
}

func stripUselessLines(itr *lexeme.Iterator) {

	newline := func(v lexeme.View) bool {
		return v.Curr().Tok == lexeme.NEWLINE
	}

	for itr.JumpToNext(newline) {

		if itr.Before() != nil && itr.Before().Tok != lexeme.NEWLINE {
			continue
		}

		if itr.After() != nil && itr.After().Tok != lexeme.NEWLINE {
			continue
		}

		itr.Remove()
	}
}

func insertSeparatorSpaces(itr *lexeme.Iterator) {

	separator := func(v lexeme.View) bool {
		return v.Curr().Tok == lexeme.DELIM
	}

	for itr.JumpToNext(separator) {
		if itr.After() != nil && itr.After().Tok != lexeme.NEWLINE {

			itr.Append(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}
	}
}

func insertCommentSpaces(itr *lexeme.Iterator) {

	comment := func(itr lexeme.View) bool {
		return itr.Curr().Tok == lexeme.COMMENT
	}

	for itr.JumpToNext(comment) {
		if itr.Before() != nil &&
			itr.Before().Tok != lexeme.NEWLINE &&
			itr.Before().Tok != lexeme.SPACE {

			itr.Prepend(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}
	}
}

func insertAssignmentSpaces(itr *lexeme.Iterator) {

	assignment := func(itr lexeme.View) bool {
		return itr.Curr().Tok == lexeme.ASSIGN
	}

	for itr.JumpToNext(assignment) {
		if itr.Before() != nil && itr.Before().Tok != lexeme.SPACE {
			itr.Prepend(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}

		if itr.After() != nil && itr.After().Tok != lexeme.SPACE {
			itr.Append(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}
	}
}

func insertOperatorSpaces(itr *lexeme.Iterator) {

	operator := func(itr lexeme.View) bool {
		return itr.Curr().Tok.IsOperator()
	}

	for itr.JumpToNext(operator) {
		if itr.Before() != nil && itr.Before().Tok != lexeme.SPACE {
			itr.Prepend(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}

		if itr.After() != nil && itr.After().Tok != lexeme.SPACE {
			itr.Append(&lexeme.Lexeme{
				Tok: lexeme.SPACE,
				Raw: " ",
			})
		}
	}
}

func unifyLineEndings(itr *lexeme.Iterator) {

	newline := func(v lexeme.View) bool {
		return v.Curr().Tok == lexeme.NEWLINE
	}

	lineEnding := "\n"

	if itr.JumpToNext(newline) {
		lineEnding = itr.Curr().Raw
	}

	for itr.JumpToNext(newline) {
		itr.Curr().Raw = lineEnding
	}

	if itr.Prev() && itr.Curr().Tok != lexeme.NEWLINE {
		itr.Append(&lexeme.Lexeme{
			Tok: lexeme.NEWLINE,
			Raw: lineEnding,
		})
	}
}

func indentLines(itr *lexeme.Iterator) {

	indent := 0
	preUndented := false

	for itr.Next() {
		switch {
		case itr.Curr().Tok.IsOpener():
			indent++

		case itr.Curr().Tok.IsCloser():
			if !preUndented {
				indent--
			}
			preUndented = false

		case itr.Curr().Tok != lexeme.NEWLINE:
		case itr.After() == nil:
		case itr.After().Tok == lexeme.NEWLINE:

		case itr.After().Tok.IsCloser():
			indent--
			preUndented = true

		case indent > 0:
			itr.Append(&lexeme.Lexeme{
				Tok:  lexeme.SPACE,
				Raw:  strings.Repeat("\t", indent),
				Line: itr.Curr().Line + 1,
			})
		}
	}
}

func updatePositions(itr *lexeme.Iterator) {

	line, col := 0, 0

	for itr.Next() {
		itr.Curr().Line = line
		itr.Curr().Col = col

		if itr.Curr().Tok == lexeme.NEWLINE {
			line++
			col = 0
		} else {
			col += len(itr.Curr().Raw)
		}
	}
}
