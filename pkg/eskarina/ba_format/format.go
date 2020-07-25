package format

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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

	for head != nil && head.Is(prop.PR_WHITESPACE) {
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

	nextIs := func(curr *lexeme.Lexeme, p prop.Prop) bool {
		return curr.Next != nil && curr.Next.Is(p)
	}

	for curr := head; curr != nil && curr.Next != nil; {

		var next *lexeme.Lexeme

		switch {
		case curr.Is(prop.PR_NEWLINE) && nextIs(curr, prop.PR_WHITESPACE):
			next = remove(curr.Next)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_NEWLINE):
			next = remove(curr)

		case curr.Is(prop.PR_SPELL) && nextIs(curr, prop.PR_WHITESPACE):
			next = remove(curr.Next)

		case curr.Is(prop.PR_OPENER) && nextIs(curr, prop.PR_WHITESPACE):
			next = remove(curr.Next)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_SEPARATOR):
			next = remove(curr)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_CLOSER):
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

		if !lex.Is(prop.PR_SEPARATOR) || lex.Next == nil {
			continue
		}

		if !lex.Next.Any(prop.PR_WHITESPACE, prop.PR_NEWLINE) {
			lex.Append(&lexeme.Lexeme{
				Props: []prop.Prop{
					prop.PR_REDUNDANT,
					prop.PR_WHITESPACE,
				},
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
		if lex.Is(prop.PR_WHITESPACE) {
			lex.Raw = " "
		}
	}

	return head
}

func trimEmptyLines(head *lexeme.Lexeme) *lexeme.Lexeme {

	if head == nil {
		return nil
	}

	if head.Is(prop.PR_NEWLINE) {
		for head.Next != nil && head.Next.Is(prop.PR_NEWLINE) {
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

	if !tail.Is(prop.PR_NEWLINE) {
		return head
	}

	for tail.Prev != nil && tail.Prev.Is(prop.PR_NEWLINE) {
		tail.Prev.Remove()
	}

	return head
}

func reduceEmptyLines(head *lexeme.Lexeme) *lexeme.Lexeme {

	var single, double bool

	for lex := head; lex != nil; {

		switch {
		case !lex.Is(prop.PR_NEWLINE):
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
		if lex.Is(prop.PR_NEWLINE) {
			lex.Raw = lineEnding
		}
	}

	tail := head
	for lex := head; lex != nil; lex = lex.Next {
		tail = lex
	}

	if tail != nil && !tail.Is(prop.PR_NEWLINE) {
		tail.Append(&lexeme.Lexeme{
			Props: []prop.Prop{
				prop.PR_TERMINATOR,
				prop.PR_NEWLINE,
			},
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
		case lex.Is(prop.PR_OPENER):
			indent++

		case lex.Is(prop.PR_CLOSER):
			indent--

		case !lex.Is(prop.PR_NEWLINE) || lex.Next == nil:

		case lex.Next.Is(prop.PR_NEWLINE):

		case lex.Next.Is(prop.PR_CLOSER):
			indent--

		case indent > 0 && lex.Is(prop.PR_NEWLINE):
			lex.Append(&lexeme.Lexeme{
				Props: []prop.Prop{prop.PR_REDUNDANT, prop.PR_WHITESPACE},
				Raw:   strings.Repeat("\t", indent),
				Line:  lex.Line + 1,
			})
		}
	}

	return head
}
