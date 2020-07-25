package format

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func FormatAll(first *lexeme.Lexeme, lineEnding string) *lexeme.Lexeme {
	return format(first, lineEnding)
}

func format(first *lexeme.Lexeme, lineEnding string) *lexeme.Lexeme {

	first = trimLeadingSpace(first)
	first = trimSpaces(first)
	first = insertSpaces(first)
	//	first = reduceSpaces(first)
	//	first = reduceEmptyLines(first)
	//first = unifyLineEndings(first, lineEnding)

	// 7: Align comments for consecutive lines with comments

	return first
}

func trimLeadingSpace(first *lexeme.Lexeme) *lexeme.Lexeme {

	for first != nil && first.Is(prop.PR_WHITESPACE) {
		next := first.Next
		first.Remove()
		first = next
	}

	return first
}

func trimSpaces(first *lexeme.Lexeme) *lexeme.Lexeme {

	remove := func(lex *lexeme.Lexeme) {

		if lex == first {
			first = lex.Next
		}

		lex.Remove()
	}

	nextIs := func(curr *lexeme.Lexeme, p prop.Prop) bool {
		return curr.Next != nil && curr.Next.Is(p)
	}

	for curr := first; curr != nil; curr = curr.Next {

		if curr.Next == nil {
			break
		}

		switch {
		case curr.Is(prop.PR_NEWLINE) && nextIs(curr, prop.PR_WHITESPACE):
			remove(curr.Next)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_NEWLINE):
			remove(curr)

		case curr.Is(prop.PR_SPELL) && nextIs(curr, prop.PR_WHITESPACE):
			remove(curr.Next)

		case curr.Is(prop.PR_OPENER) && nextIs(curr, prop.PR_WHITESPACE):
			remove(curr.Next)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_SEPARATOR):
			remove(curr)

		case curr.Is(prop.PR_WHITESPACE) && nextIs(curr, prop.PR_CLOSER):
			remove(curr)
		}
	}

	return first
}

func insertSpaces(first *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := first; lex != nil; lex = lex.Next {

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

	return first
}

func reduceSpaces(first *lexeme.Lexeme) *lexeme.Lexeme {

	for lex := first; lex != nil; lex = lex.Next {
		if lex.Is(prop.PR_WHITESPACE) {
			lex.Raw = " "
		}
	}

	return first
}

/*
func reduceEmptyLines(in, out chan token.Token) {

	var single, double bool

	for tk := range in {

		switch {
		case !tk.Is(PR_NEWLINE):
			single, double = false, false

		case !single:
			single = true

		case !double:
			double = true

		default:
			continue
		}

		out <- tk
	}

	close(out)
}

func unifyLineEndings(in, out chan token.Token, lineEnding string) {

	for tk := range in {
		if tk.Is(PR_NEWLINE) && tk.Raw() != lineEnding {
			tk = newlineToken(tk, lineEnding)
		}

		out <- tk
	}

	close(out)
}

func newlineToken(curr token.Token, ending string) token.Token {

	new := token.Tok{
		RawProps: curr.Props(),
		RawStr:   ending,
	}

	new.Line, new.ColBegin = curr.Begin()
	_, new.ColEnd = curr.End()

	return new
}
*/
