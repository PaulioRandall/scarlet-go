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

// SnippetStream provides access to an ordered stream of snippets.
type SnippetStream interface {

	// Read returns the next Snippet in the stream. An EOF snippet is always
	// returned if the stream is empty. Within Scarlet, an expression is only a
	// statement if it does not form the immediate tokens of another statement.
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
	var snip Snippet
	ts := token.ToStream(tks)
	ss := impl{ts}

	for snip = ss.Read(); snip.Kind != SNIPPET_EOF; snip = ss.Read() {
		snippets = append(snippets, snip)
	}

	return append(snippets, snip)
}

// PrintAll pretty prints all snippets in snips.
func PrintAll(snips []Snippet) {
	printSnippets(snips, 0)
}

func printSnippets(snips []Snippet, indent int) {
	for _, snip := range snips {
		if snip.Kind == SNIPPET_EOF {
			println(SNIPPET_EOF)
			println()
		} else {
			printSnippet(snip, indent)
		}
	}
}

func printSnippet(snip Snippet, indent int) {

	printIndent(indent)

	for i, tk := range snip.Tokens {
		if i != 0 {
			print(" ")
		}

		print(tk.Lexeme)
	}

	println()
	printSnippets(snip.Snippets, indent+1)
}

func printIndent(indent int) {
	s := strings.Repeat("  ", indent)
	print(s)
}
