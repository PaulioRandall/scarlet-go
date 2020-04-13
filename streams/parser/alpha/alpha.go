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

func PartitionAll(tks []lexeme.Token) []AlphaStatement {
	itr := lexeme.NewIterator(tks)
	return partitionStatements(itr)
}

func partitionStatements(itr *lexeme.TokenIterator) []AlphaStatement {

	var as []AlphaStatement

	for tk := itr.Peek(); tk.Lexeme != lexeme.LEXEME_EOF; tk = itr.Peek() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR {
			_ = itr.Next()
			continue
		}

		expectNotEmpty(tk, itr)
		s := partitionStatement(itr)
		as = append(as, s)
	}

	return as
}

func partitionStatement(itr *lexeme.TokenIterator) AlphaStatement {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR
	var a AlphaStatement

	for tk := itr.Next(); tk.Lexeme != TERMINATOR; tk = itr.Next() {

		expectNotEmpty(tk, itr)

		a.Tokens = append(a.Tokens, tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			a.Subs = partitionBlock(itr)
		}
	}

	return a
}

func partitionBlock(itr *lexeme.TokenIterator) []AlphaStatement {

	var as []AlphaStatement
	var tk lexeme.Token

	for tk = itr.Next(); tk.Lexeme != lexeme.LEXEME_END; tk = itr.Next() {

		expectNotEmpty(tk, itr)

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			// Duplicate TERMINATORS should have been removed by this stage
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		a := partitionStatement(itr)
		as = append(as, a)
	}

	return as
}

func expectNotEmpty(tk lexeme.Token, itr *lexeme.TokenIterator) {
	if tk == (lexeme.Token{}) {
		panic(newErr(tk, `[Partitioner] Unexpected empty token at index %d`, itr.Index()))
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
