package format

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/prop"
	"github.com/PaulioRandall/scarlet-go/pkg/esmerelda/shared/token"
)

func FormatAll(tks []token.Token, lineEnding string) []token.Token {

	in := make(chan token.Token)
	out := make(chan token.Token)

	go func() {
		for _, tk := range tks {
			in <- tk
		}
		close(in)
	}()

	go Format(in, out, lineEnding)

	r := []token.Token{}
	for tk := range out {
		r = append(r, tk)
	}

	return r
}

func Format(in, out chan token.Token, lineEnding string) {

	trimmed := make(chan token.Token)
	go trimSpaces(in, trimmed)

	reduced := make(chan token.Token)
	go reduceSpaces(trimmed, reduced)

	aligned := make(chan token.Token)
	go reduceEmptyLines(reduced, aligned)

	unified := make(chan token.Token)
	go unifyLineEndings(aligned, unified, lineEnding)

	for tk := range unified {
		out <- tk
	}
	close(out)
	// Use a functional approach:
	// 7: Align comments for consecutive lines with comments
}

func trimSpaces(in, out chan token.Token) {

	a := <-in
	for a.Is(PR_WHITESPACE) {
		a = <-in
	}

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
