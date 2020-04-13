// snippet package was created to separate the concern of grouping tokens into
// statements from the parsing of those statements. The API uses a TokenStream
// create a SnippetStream which is uses the production rules to produce
// Snippets; a group of tokens and sub-snippets.
//
// Key decisions: N/A
//
// This package does not identify expressions or parse the statements. It just
// groups tokens and sub-snippets together so further processing is easier.
package statement

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// statement represents an unparsed statement, perhaps containing
// sub-statements.
type Statement struct {
	Tokens []lexeme.Token
	Stats  []Statement
}

// StatementStream provides access to an ordered stream of unparsed statements.
type StatementStream interface {

	// Read returns the next Statement in the stream. If the stream is empty a
	// statement with both values a nil will be returned. Within Scarlet, an
	// expression is only part of a statement if it does not form the immediate
	// tokens of a sub-statement.
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
	Read() Statement
}

// PartitionAll consumes all tokens from ts, runs them through a
// StatementStream, then returns the resultant statements as an array.
func PartitionAll(tks []lexeme.Token) []Statement {

	var stats []Statement
	ts := token.ToStream(tks)
	ss := impl{ts, nil}

	for stat := ss.Read(); stat.Tokens != nil; stat = ss.Read() {
		stats = append(stats, stat)
	}

	return stats
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
