package format

import (
	//"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func FormatAll(tks []token.Token) []token.Token {

	in := make(chan token.Token)
	out := make(chan token.Token)

	go func() {
		for _, tk := range tks {
			in <- tk
		}
		close(in)
	}()

	Format(in, out)

	tks = []token.Token{}
	for tk := range out {
		tks = append(tks, tk)
	}

	return tks
}

func Format(in, out chan token.Token) {

	var lineEnding string

	searched := make(chan token.Token)
	go detectLineEndings(in, searched, &lineEnding)

	trimmed := make(chan token.Token)
	go trimSpaces(searched, trimmed)

	reduced := make(chan token.Token)
	go reduceSpaces(trimmed, reduced)

	aligned := make(chan token.Token)
	go reduceEmptyLines(reduced, aligned)

	// https://en.wikipedia.org/wiki/Parsing_expression_grammar

	// Use a functional approach:
	// 2: Remove all redundant whitespace
	// 3: Remove multiple empty lines
	// 4: Insert single space after value separators if not a newline, i.e. ','
	// 5: Remove empty lines between list items
	// 6: Indent for multiline statements (except initiating line and final ')')
	// 7: Align comments for consecutive lines with comments

}

func detectLineEndings(in, out chan token.Token, ending *string) {

	var found bool

	for tk := range in {
		out <- tk

		if !found && tk.Is(PR_NEWLINE) {
			found = true
			*ending = tk.Raw()
		}
	}
}

func trimSpaces(in, out chan token.Token) {

	a := <-in

	for a != nil {
		b := <-in

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
		case a.Is(PR_WHITESPACE) && b.Is(PR_CLOSER):
			a = b

		default:
			out <- a
			a = b
		}
	}

	close(out)
}

func reduceSpaces(in, out chan token.Token) {

	for tk := range in {
		if tk.Is(PR_WHITESPACE) && tk.Raw() != " " {
			tk = newSpace(tk)
		}

		out <- tk
	}

	close(out)
}

func newSpace(curr token.Token) token.Token {

	new := token.Tok{
		RawProps: curr.Props(),
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
