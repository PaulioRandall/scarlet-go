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
	//first = trimSpaces(first)
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

/*
func trimSpaces(first *lexeme.Lexeme) *lexeme.Lexeme {


	for lex := first; lex != nil; lex = lex.Next {

		if b == nil {
			out <- a
			break
		}

		switch {
		case a.Is(PR_NEWLINE) && b.Is(PR_WHITESPACE):
		case a.Is(PR_WHITESPACE) && b.Is(PR_NEWLINE):
			a = b
		case a.Is(PR_SPELL) && b.Is(PR_WHITESPACE):
		case a.Is(PR_OPENER) && b.Is(PR_WHITESPACE):
		case a.Is(PR_WHITESPACE) && b.Is(PR_SEPARATOR):
			a = b
		case a.Is(PR_SEPARATOR) && !b.Any(PR_WHITESPACE, PR_NEWLINE):
			out <- a
			out <- newSpaceToken(b)
			a = b
		case a.Is(PR_WHITESPACE) && b.Is(PR_CLOSER):
			a = b

		default:
			out <- a
			a = b
		}
	}

	close(out)
}
/*
func reduceSpaces(in, out chan token.Token) {

	for tk := range in {
		if tk.Is(PR_WHITESPACE) && tk.Raw() != " " {
			tk = newSpaceToken(tk)
		}

		out <- tk
	}

	close(out)
}

func newSpaceToken(curr token.Token) token.Token {

	new := token.Tok{
		RawProps: []Prop{PR_WHITESPACE},
		RawStr:   " ",
	}

	new.Line, new.ColBegin = curr.Begin()
	_, new.ColEnd = curr.End()

	return new
}

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
