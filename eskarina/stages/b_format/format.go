package format

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type Iterator interface {
	//AsIterator() *Iterator
	ToContainer() *lexeme.Container
	//HasPrev() bool
	//HasNext() bool
	Prev() bool
	Next() bool
	Curr() *lexeme.Lexeme
	Remove() *lexeme.Lexeme
	//Prepend(*Lexeme)
	Append(*lexeme.Lexeme)
	Before() *lexeme.Lexeme
	After() *lexeme.Lexeme
	//String() string
}

func FormatAll(con *lexeme.Container) *lexeme.Container {
	return format(con)
}

func format(con *lexeme.Container) *lexeme.Container {

	con = trimWhiteSpace(con)
	con = stripUselessLines(con)
	con = insertWhiteSpace(con)
	con = unifyLineEndings(con)
	con = indentLines(con)
	con = updatePositions(con)
	con = alignComments(con)

	return con
}

func trimWhiteSpace(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	for itr.Next() {
		if itr.Curr().Tok == lexeme.WHITESPACE {
			itr.Remove()
		}
	}

	return itr.ToContainer()
}

func stripUselessLines(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	for itr.Next() {

		if itr.Curr().Tok != lexeme.NEWLINE {
			continue
		}

		if itr.Before() == nil || itr.Before().Tok == lexeme.NEWLINE {
			if itr.After() == nil || itr.After().Tok == lexeme.NEWLINE {
				itr.Remove()
			}
		}
	}

	return itr.ToContainer()
}

func insertWhiteSpace(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	for itr.Next() {

		switch {
		case itr.After() == nil:
		case itr.After().Tok == lexeme.NEWLINE:

		case itr.Curr().Tok == lexeme.SEPARATOR:
			itr.Append(&lexeme.Lexeme{
				Tok:  lexeme.WHITESPACE,
				Raw:  " ",
				Line: itr.Curr().Line,
				Col:  itr.Curr().Col + 1,
			})
		}
	}

	return itr.ToContainer()
}

func unifyLineEndings(con *lexeme.Container) *lexeme.Container {

	lineEnding := "\n"
	itr := Iterator(con.ToIterator())

	for itr.Next() {
		if itr.Curr().Tok == lexeme.NEWLINE {
			lineEnding = itr.Curr().Raw
			break
		}
	}

	for itr.Next() {
		if itr.Curr().Tok == lexeme.NEWLINE {
			itr.Curr().Raw = lineEnding
		}
	}

	if itr.Prev() && itr.Curr().Tok != lexeme.NEWLINE {
		itr.Append(&lexeme.Lexeme{
			Tok:  lexeme.NEWLINE,
			Raw:  lineEnding,
			Line: itr.Curr().Line,
			Col:  itr.Curr().Col + len(itr.Curr().Raw),
		})
	}

	return itr.ToContainer()
}

func indentLines(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())
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

	itr := Iterator(con.ToIterator())
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

func alignComments(con *lexeme.Container) *lexeme.Container {

	itr := Iterator(con.ToIterator())

	// 1. Split into lines
	// 2. Mark lines with a comment but not ones that start with the comment
	// 3. Find consecutive comment lines (comment group)
	// 		Find the comment with the greatest col index
	//    Insert whitespace before comments in each line to match the greatest
	// 4. Join lines back together

	for itr.Next() {

	}

	return itr.ToContainer()
}
