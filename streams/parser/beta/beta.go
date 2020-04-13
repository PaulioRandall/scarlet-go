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
	Arts   []BetaStatement
}

// ArticulateAll consumes all statements from stats, runs them through a
// ArticulateStream, then returns the resultant articulates as an array.
func ArticulateAll(stats []alpha.AlphaStatement) []BetaStatement {
	return articulateStatments(stats)
}

func articulateStatments(stats []alpha.AlphaStatement) []BetaStatement {

	var arts []BetaStatement
	itr := statItr{stats, len(stats), 0}

	for stat, ok := itr.next(); ok; stat, ok = itr.next() {
		a := articulateStatment(stat)
		arts = append(arts, a)
	}

	return arts
}

func articulateStatment(stat alpha.AlphaStatement) BetaStatement {

	var art BetaStatement
	tks := stat.Tokens
	i := indexOfAssignment(tks)

	if i != -1 {
		art.Assign = tks[i]
		art.IDs = tks[:i]
		art.Exprs = tks[i+1:]
	} else {
		art.Exprs = tks
	}

	art.Arts = articulateStatments(stat.Stats)

	return art
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

// PrintAll pretty prints all BetaStatement in arts.
func PrintAll(arts []BetaStatement) {
	printBetaStatements(arts, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printBetaStatement prints all BetaStatement in arts indenting all output to the
// specified level.
func printBetaStatements(arts []BetaStatement, indent int) {
	for _, a := range arts {
		printBetaStatement(a, indent)
	}
}

// printBetaStatement prints a indenting all output to the specified level.
func printBetaStatement(a BetaStatement, indent int) {

	print(strings.Repeat("  ", indent))

	if a.Assign != (lexeme.Token{}) {
		printTokens(a.IDs)
		print(" " + a.Assign.Lexeme + " ")
	}

	printTokens(a.Exprs)

	println()
	printBetaStatements(a.Arts, indent+1)
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
