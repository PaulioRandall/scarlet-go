package format

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func FormatAll(con *lexeme.Container) *lexeme.Container {
	return format(con)
}

func format(con *lexeme.Container) *lexeme.Container {

	itr := con.ToIterator()

	trimWhiteSpace(itr)

	itr.Restart()
	stripUselessLines(itr)

	itr.Restart()
	insertSeparatorSpaces(itr)

	itr.Restart()
	unifyLineEndings(itr)

	con = itr.ToContainer()

	con = indentLines(con)
	con = updatePositions(con)
	//con = alignComments(con)

	return con
}

func trimWhiteSpace(itr *lexeme.Iterator) {

	whitespace := func(it lexeme.Iterator) bool {
		return it.Curr().Tok == lexeme.WHITESPACE
	}

	for itr.JumpToNext(whitespace) {
		itr.Remove()
	}
}

func stripUselessLines(itr *lexeme.Iterator) {

	newline := func(it lexeme.Iterator) bool {
		return it.Curr().Tok == lexeme.NEWLINE
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

	separator := func(it lexeme.Iterator) bool {
		return it.Curr().Tok == lexeme.SEPARATOR
	}

	for itr.JumpToNext(separator) {
		if itr.After() != nil && itr.After().Tok != lexeme.NEWLINE {

			itr.Append(&lexeme.Lexeme{
				Tok:  lexeme.WHITESPACE,
				Raw:  " ",
				Line: itr.Curr().Line,
				Col:  itr.Curr().Col + 1,
			})
		}
	}
}

func unifyLineEndings(itr *lexeme.Iterator) {

	newline := func(it lexeme.Iterator) bool {
		return it.Curr().Tok == lexeme.NEWLINE
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

func indentLines(con *lexeme.Container) *lexeme.Container {

	itr := con.ToIterator()
	indent := 0

	for itr.Next() {
		switch {
		case itr.Curr().Tok.IsOpener():
			indent++

		case itr.Curr().Tok.IsCloser():
			indent--

		case itr.Curr().Tok != lexeme.NEWLINE:
		case itr.After() == nil:
		case itr.After().Tok == lexeme.NEWLINE:

		case itr.After().Tok.IsCloser():
			indent--

		case indent > 0:
			itr.Append(&lexeme.Lexeme{
				Tok:  lexeme.WHITESPACE,
				Raw:  strings.Repeat("\t", indent),
				Line: itr.Curr().Line + 1,
			})
		}
	}

	return itr.ToContainer()
}

func updatePositions(con *lexeme.Container) *lexeme.Container {

	itr := con.ToIterator()
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

	return itr.ToContainer()
}
