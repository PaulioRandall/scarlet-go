// snippet package was created to separate the concern of grouping tokens into
// statements from the parsing of those statements. The API uses a TokenStream
// create a SnippetStream which is uses the production rules to produce
// Snippets; a group of tokens and sub-snippets.
//
// Key decisions: N/A
//
// This package does not identify expressions or parse the statements. It just
// groups tokens and sub-snippets together so further processing is easier.
package snippet

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// Snippet represents an unparsed statement, perhaps containing sub-statements.
type Snippet struct {
	Tokens   []lexeme.Token
	Snippets []Snippet
}

// SnippetStream provides access to an ordered stream of snippets.
type SnippetStream interface {

	// Read returns the next Snippet in the stream. If the stream is empty a
	// snippet with both values a nil will be returned. Within Scarlet, an
	// expression is only a statement if it does not form the immediate tokens of
	// another statement.
	//
	// E.g.
	// Consider `x := 1 + 1`:
	// - the whole thing is a statement, a snippet
	// - `x := 1 + 1` will all become tokens in the snippet
	// - `1 + 1` is an expression which the SnippetStream nor Snippet know about
	// Now consider `f := F(a, b -> c) c := a + b`:
	// - the whole thing is a statement, a snippet
	// - `f := F(a, b -> c)` will all become the snippet tokens
	// - `c := a + b` will become a sub-snippet
	Read() Snippet
}

// GroupAll consumes all tokens from ts, runs them through a SnippetStream, then
// returns the resultant snippets as an array.
func GroupAll(tks []lexeme.Token) []Snippet {

	var snippets []Snippet
	ts := token.ToStream(tks)
	ss := impl{ts, nil}

	for snip := ss.Read(); snip.Tokens != nil; snip = ss.Read() {
		snippets = append(snippets, snip)
	}

	return snippets
}

// PrintAll pretty prints all snippets in s.
func PrintAll(s []Snippet) {
	printSnippets(s, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printSnippet prints all snippets in s indenting all output to the specified
// level.
func printSnippets(s []Snippet, indent int) {
	for _, snip := range s {
		printSnippet(snip, indent)
	}
}

// printSnippet prints the s indenting all output to the specified level.
func printSnippet(s Snippet, indent int) {

	print(strings.Repeat("  ", indent))

	for i, tk := range s.Tokens {
		if i != 0 {
			print(" ")
		}

		print(tk.Lexeme)
	}

	println()
	printSnippets(s.Snippets, indent+1)
}
