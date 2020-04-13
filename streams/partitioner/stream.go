// partitioner package was created to separate the concern of grouping tokens into
// statements from the parsing of those statements. The API uses a TokenStream
// create a SnippetStream which is uses the production rules to produce
// Snippets; a group of tokens and sub-snippets.
//
// Key decisions: N/A
//
// This package does not identify expressions or parse the statements. It just
// groups tokens and sub-statement together so further processing is easier.
package partitioner

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// statement represents an unparsed statement, perhaps containing
// sub-statements.
//
// E.g.
// Consider `x := 1 + 1`:
// - the whole thing is a statement
// - `x := 1 + 1` will all become tokens in the statement
// - `1 + 1` is an expression which the StatementStream nor Statement know
// about. Now consider `f := F(a, b -> c) c := a + b`:
// - the whole thing is a statement
// - `f := F(a, b -> c)` will all become the statement tokens
// - `c := a + b` will become a sub-statement
type Statement struct {
	Tokens []lexeme.Token
	Stats  []Statement
}

func PartitionAll(tks []lexeme.Token) []Statement {
	itr := lexeme.NewIterator(tks)
	return partitionStatements(itr)
}

func partitionStatements(itr *lexeme.TokenIterator) []Statement {

	var stats []Statement

	for tk := itr.Peek(); tk.Lexeme != lexeme.LEXEME_EOF; tk = itr.Peek() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR {
			_ = itr.Next()
			continue
		}

		expectNotEmpty(tk, itr)
		s := partitionStatement(itr)
		stats = append(stats, s)
	}

	return stats
}

func partitionStatement(itr *lexeme.TokenIterator) Statement {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR
	var stat Statement

	for tk := itr.Next(); tk.Lexeme != TERMINATOR; tk = itr.Next() {

		expectNotEmpty(tk, itr)

		stat.Tokens = append(stat.Tokens, tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			stat.Stats = partitionBlock(itr)
		}
	}

	return stat
}

func partitionBlock(itr *lexeme.TokenIterator) []Statement {

	var stats []Statement

	for tk := itr.Next(); tk.Lexeme != lexeme.LEXEME_END; tk = itr.Next() {

		expectNotEmpty(tk, itr)

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			// Duplicate TERMINATORS should have been removed by this stage
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		s := partitionStatement(itr)
		stats = append(stats, s)
	}

	return stats
}

func expectNotEmpty(tk lexeme.Token, itr *lexeme.TokenIterator) {
	if tk == (lexeme.Token{}) {
		panic(newErr(tk, `[Partitioner] Unexpected empty token at index %d`, itr.Index()))
	}
}

// PrintAll pretty prints all statement in s.
func PrintAll(s []Statement) {
	printStatements(s, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printStatements prints all statements in stats indenting all output to the
// specified level.
func printStatements(stats []Statement, indent int) {
	for _, s := range stats {
		printStatement(s, indent)
	}
}

// printStatement prints s indenting all output to the specified level.
func printStatement(s Statement, indent int) {

	print(strings.Repeat("  ", indent))

	for i, tk := range s.Tokens {
		if i != 0 {
			print(" ")
		}

		print(tk.Lexeme)
	}

	println()
	printStatements(s.Stats, indent+1)
}
