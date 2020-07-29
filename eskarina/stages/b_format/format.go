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
	//Prev() bool
	Next() bool
	Curr() *lexeme.Lexeme
	Remove() *lexeme.Lexeme
	//Prepend(*Lexeme)
	Append(*lexeme.Lexeme)
	//Before() *Lexeme
	After() *lexeme.Lexeme
	//String() string
}

func FormatAll(con *lexeme.Container, lineEnding string) *lexeme.Container {
	return format(con, lineEnding)
}

func format(con *lexeme.Container, lineEnding string) *lexeme.Container {

	con = trimWhiteSpace(con)
	con = insertWhiteSpace(con)
	//head = reduceSpaces(head)
	//head = trimEmptyLines(head)
	//head = reduceEmptyLines(head)
	//head = unifyLineEndings(head, lineEnding)
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

/*
func reduceSpaces(head *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := head; lex != nil; lex = lex.Next {
		if lex.Tok == lexeme.WHITESPACE {
			lex.Raw = " "
		}
	}

	return head
}
/*
func trimEmptyLines(head *lexeme.Lexeme) *lexeme.Lexeme {

	if head == nil {
		return nil
	}

	if head.Tok == lexeme.NEWLINE {
		for head.Next != nil && head.Next.Tok == lexeme.NEWLINE {
			head.Next.Remove()
		}
	}

	if head == nil {
		return nil
	}

	tail := head
	for lex := head; lex != nil; lex = lex.Next {
		tail = lex
	}

	if tail.Tok != lexeme.NEWLINE {
		return head
	}

	for tail.Prev != nil && tail.Prev.Tok == lexeme.NEWLINE {
		tail.Prev.Remove()
	}

	return head
}
/*
func reduceEmptyLines(head *lexeme.Lexeme) *lexeme.Lexeme {

	var single, double bool

	for lex := head; lex != nil; {

		switch {
		case lex.Tok != lexeme.NEWLINE:
			single, double = false, false

		case !single:
			single = true

		case !double:
			double = true

		default:
			next := lex.Next
			lex.Remove()
			lex = next
			continue
		}

		lex = lex.Next
	}

	return head
}
/*
func unifyLineEndings(head *lexeme.Lexeme, lineEnding string) *lexeme.Lexeme {

	for lex := head; lex != nil; lex = lex.Next {
		if lex.Tok == lexeme.NEWLINE {
			lex.Raw = lineEnding
		}
	}

	tail := head
	for lex := head; lex != nil; lex = lex.Next {
		tail = lex
	}

	if tail != nil && tail.Tok != lexeme.NEWLINE {
		tail.Append(&lexeme.Lexeme{
			Raw:  lineEnding,
			Line: tail.Line,
		})
	}

	return head
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
