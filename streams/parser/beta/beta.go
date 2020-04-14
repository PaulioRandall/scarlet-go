// beta package was created to split statements into separate assignment
// and expression parts. The API uses a StatementStream to create an
// ArticulateStream which uses the production rules to produce Articulates; a
// statement with an identified type (Kind) with separate stores for assignment
// targets and source expressions.
//
// Key decisions: N/A
//
// This package does not identify expressions or parse the statement parts. It
// just identifies splits the assignment targets from their sources for easier
// processing down the line.
package beta

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/parser/alpha"
)

// BetaStatement represents an unparsed statement with separate members for
// storing assignment token (Assign), it it has one, the assignment targets
// (IDs), if it has any, and any expressions (Exprs).
//
// E.g.
// Consider `x := 1 + 1`:
// - the whole thing is a statement
// - `:=` will become the assignment token (Assign)
// - `x` will become the identifier tokens (IDs)
// - `1 + 1` will become the expression tokens (Exprs)
//
// Now consider `a, b := 1 + 2, 4 - 3`:
// - the whole thing is a statement
// - `:=` will become the assignment token (Assign)
// - `a, b` will become the identifier tokens (IDs)
// - `1 + 2, 4 - 3` will all become the expression tokens (Exprs)
type BetaStatement struct {
	Assign lexeme.Token
	IDs    []lexeme.Token
	Exprs  []lexeme.Token
	Subs   []BetaStatement
}

// TransformAll consumes all statements from as, runs them through a
// ArticulateStream, then returns the resultant articulates as an array.
func TransformAll(as []alpha.AlphaStatement) []BetaStatement {
	return transformStatments(as)
}

func transformStatments(as []alpha.AlphaStatement) []BetaStatement {

	var bs []BetaStatement
	itr := statItr{as, len(as), 0}

	for a, ok := itr.next(); ok; a, ok = itr.next() {
		b := transformStatment(a)
		bs = append(bs, b)
	}

	return bs
}

func transformStatment(a alpha.AlphaStatement) BetaStatement {

	var b BetaStatement
	tks := a.Tokens
	i := indexOfAssignment(tks)

	if i != -1 {
		b.Assign = tks[i]
		b.IDs = tks[:i]
		b.Exprs = tks[i+1:]
	} else {
		b.Exprs = tks
	}

	b.Subs = transformStatments(a.Subs)

	return b
}

func indexOfAssignment(tks []lexeme.Token) int {
	for i, tk := range tks {

		switch tk.Lexeme {
		case lexeme.LEXEME_ID, lexeme.LEXEME_DELIM, lexeme.LEXEME_VOID:
			continue
		case lexeme.LEXEME_ASSIGN:
			return i
		}

		break
	}

	return -1
}

// PrintAll pretty prints all BetaStatement in bs.
func PrintAll(bs []BetaStatement) {
	printBetaStatements(bs, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printBetaStatement prints all BetaStatement in bs indenting all output to the
// specified level.
func printBetaStatements(bs []BetaStatement, indent int) {
	for _, b := range bs {
		printBetaStatement(b, indent)
	}
}

// printBetaStatement prints b indenting all output to the specified level.
func printBetaStatement(b BetaStatement, indent int) {

	print(strings.Repeat("  ", indent))

	if b.Assign != (lexeme.Token{}) {
		printTokens(b.IDs)
		print(" " + b.Assign.Lexeme + " ")
	}

	printTokens(b.Exprs)

	println()
	printBetaStatements(b.Subs, indent+1)
}

// printTokens prints a slice of tokens.
func printTokens(tks []lexeme.Token) {

	print("[")

	for i, tk := range tks {
		if i != 0 {
			print(" ")
		}

		print(tk.Lexeme)
	}

	print("]")
}
