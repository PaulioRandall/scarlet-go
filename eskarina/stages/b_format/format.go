package format

/*
import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

func FormatAll(head *lexeme.Lexeme, lineEnding string) *lexeme.Lexeme {
	return format(head, lineEnding)
}

func format(head *lexeme.Lexeme, lineEnding string) *lexeme.Lexeme {

	head = trimLeadingSpace(head)
	head = trimSpaces(head)
	head = insertSpaces(head)
	head = reduceSpaces(head)
	head = trimEmptyLines(head)
	head = reduceEmptyLines(head)
	head = unifyLineEndings(head, lineEnding)
	head = indentNests(head)
	//head = alignComments(head)

	return head
}

func trimLeadingSpace(head *lexeme.Lexeme) *lexeme.Lexeme {

	for head != nil && head.Tok == lexeme.WHITESPACE {
		next := head.Next
		head.Remove()
		head = next
	}

	return head
}

func trimSpaces(head *lexeme.Lexeme) *lexeme.Lexeme {

	remove := func(lex *lexeme.Lexeme) *lexeme.Lexeme {

		if lex == head {
			head = lex.Next
		}

		next := lex.Next
		lex.Remove()
		return next
	}

	nextTok := func(curr *lexeme.Lexeme) lexeme.Token {
		if curr.Next == nil {
			return lexeme.UNDEFINED
		}
		return curr.Next.Tok
	}

	for curr := head; curr != nil && curr.Next != nil; {

		var next *lexeme.Lexeme

		switch {
		case curr.Tok == lexeme.NEWLINE && nextTok(curr) == lexeme.WHITESPACE:
			next = remove(curr.Next)

		case curr.Tok == lexeme.WHITESPACE && nextTok(curr) == lexeme.NEWLINE:
			next = remove(curr)

		case curr.Tok == lexeme.SPELL && nextTok(curr) == lexeme.WHITESPACE:
			next = remove(curr.Next)

		case curr.Tok.IsOpener() && nextTok(curr) == lexeme.WHITESPACE:
			next = remove(curr.Next)

		case curr.Tok == lexeme.WHITESPACE && nextTok(curr) == lexeme.SEPARATOR:
			next = remove(curr)

		case curr.Tok == lexeme.WHITESPACE && nextTok(curr).IsCloser():
			next = remove(curr)

		default:
			next = curr.Next
		}

		curr = next
	}

	return head
}

func insertSpaces(head *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := head; lex != nil; lex = lex.Next {

		if lex.Tok != lexeme.SEPARATOR || lex.Next == nil {
			continue
		}

		if !lex.Next.Tok.IsAny(lexeme.WHITESPACE, lexeme.NEWLINE) {
			lex.Append(&lexeme.Lexeme{
				Tok:  lexeme.WHITESPACE,
				Raw:  " ",
				Line: lex.Line,
				Col:  lex.Col + 1,
			})
		}
	}

	return head
}

func reduceSpaces(head *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := head; lex != nil; lex = lex.Next {
		if lex.Tok == lexeme.WHITESPACE {
			lex.Raw = " "
		}
	}

	return head
}

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

func alignComments(head *lexeme.Lexeme) *lexeme.Lexeme {

	return head
}
*/
