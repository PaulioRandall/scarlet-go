package format

import (
	//"strings"

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
	//head = indentNests(head)
	//head = alignComments(head)

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

/*
func indentNests(head *lexeme.Lexeme) *lexeme.Lexeme {

	indent := 0

	for lex := head; lex != nil; lex = lex.Next {

		switch {
		case lex.Tok.IsOpener():
			indent++

		case lex.Tok.IsCloser():
			indent--

		case lex.Tok != lexeme.NEWLINE || lex.Next == nil:

		case lex.Next.Tok == lexeme.NEWLINE:

		case lex.Next.Tok.IsCloser():
			indent--

		case indent > 0 && lex.Tok == lexeme.NEWLINE:
			lex.Append(&lexeme.Lexeme{
				Tok:  lexeme.WHITESPACE,
				Raw:  strings.Repeat("\t", indent),
				Line: lex.Line + 1,
			})
		}
	}

	return head
}
/*
func alignComments(head *lexeme.Lexeme) *lexeme.Lexeme {

	return head
}
*/
