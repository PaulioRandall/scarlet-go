// alpha package was created to separate the concern of grouping tokens into
// statements from the parsing of those statements.
//
// Key decisions: N/A
//
// This package does not identify expressions or parse the statements. It just
// groups tokens and sub-statement together so further processing is easier.
package alpha

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// AlphaStatement represents an unparsed statement, perhaps containing
// sub-statements.
//
// E.g.
// Consider `x := 1 + 1`:
// - the whole thing is a statement
// - `x := 1 + 1` will all become tokens in the statement
// - `1 + 1` is an expression which the AlphaStatement knows nothing about.
// Now consider `f := F(a, b -> c) c := a + b`:
// - the whole thing is a statement
// - `f := F(a, b -> c)` will all become the statement tokens
// - `c := a + b` will become a sub-statement
type AlphaStatement struct {
	Tokens []lexeme.Token
	Subs   []AlphaStatement
}

func TransformAll(tks []lexeme.Token) []AlphaStatement {
	itr := lexeme.NewIterator(tks)
	return transformStatements(itr)
}

func transformStatements(itr *lexeme.TokenIterator) []AlphaStatement {

	var as []AlphaStatement

	for tk := itr.Peek(); tk.Lexeme != lexeme.LEXEME_EOF; tk = itr.Peek() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR {
			itr.Skip()
			continue
		}

		expectNot(itr, tk, lexeme.LEXEME_UNDEFINED)

		s := transformStatement(itr)
		as = append(as, s)
	}

	return as
}

func transformStatement(itr *lexeme.TokenIterator) AlphaStatement {

	var a AlphaStatement

	for tk := itr.Peek(); tk.Lexeme != lexeme.LEXEME_TERMINATOR; tk = itr.Peek() {

		expectNot(itr, tk, lexeme.LEXEME_UNDEFINED)
		expectNot(itr, tk, lexeme.LEXEME_EOF)

		a.Tokens = append(a.Tokens, itr.Next())

		if tk.Lexeme == lexeme.LEXEME_DO {
			a.Subs = transformBlock(itr)
		}
	}

	expect(itr, itr.Next(), lexeme.LEXEME_TERMINATOR)
	return a
}

func transformBlock(itr *lexeme.TokenIterator) []AlphaStatement {

	var as []AlphaStatement

	for tk := itr.Peek(); tk.Lexeme != lexeme.LEXEME_END; tk = itr.Peek() {

		expectNot(itr, tk, lexeme.LEXEME_UNDEFINED)
		expectNot(itr, tk, lexeme.LEXEME_TERMINATOR)
		expectNot(itr, tk, lexeme.LEXEME_EOF)

		a := transformStatement(itr)
		as = append(as, a)
	}

	return as
}

func expect(itr *lexeme.TokenIterator, tk lexeme.Token, lex lexeme.Lexeme) {
	if tk.Lexeme != lex {
		panic(newErr(tk, `[alpha] Expected %s, but got %s at index %d`, lex, tk.Lexeme, itr.Index()))
	}
}

func expectNot(itr *lexeme.TokenIterator, tk lexeme.Token, lex lexeme.Lexeme) {
	if tk.Lexeme == lex {
		panic(newErr(tk, `[alpha] Unexpected %s at index %d`, lex, itr.Index()))
	}
}

// PrintAll pretty prints all statement in as.
func PrintAll(as []AlphaStatement) {
	printStatements(as, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printStatements prints all statements in as indenting all output to the
// specified level.
func printStatements(as []AlphaStatement, indent int) {
	for _, a := range as {
		printStatement(a, indent)
	}
}

// printStatement prints a indenting all output to the specified level.
func printStatement(a AlphaStatement, indent int) {

	print(strings.Repeat("  ", indent))

	for i, tk := range a.Tokens {
		if i != 0 {
			print(" ")
		}

		print(tk.Lexeme)
	}

	println()
	printStatements(a.Subs, indent+1)
}
