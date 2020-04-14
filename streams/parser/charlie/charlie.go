// charlie package parses the assignment tokens within BetaStatements to
// produce CharlieStatements.
package charlie

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/parser/beta"
)

// CharlieStatement represents an partially parsed statement where the
// assignment tokens have been parsed but the expressions tokens have not.
type CharlieStatement struct {
	Assign lexeme.Token
	IDs    []lexeme.Token
	Exprs  []lexeme.Token
	Subs   []CharlieStatement
}

func TransformAll(bs []beta.BetaStatement) []CharlieStatement {
	return transformStatments(bs)
}

func transformStatments(bs []beta.BetaStatement) []CharlieStatement {

	var cs []CharlieStatement
	itr := statItr{bs, len(bs), 0}

	for b, ok := itr.next(); ok; b, ok = itr.next() {
		c := transformStatment(b)
		cs = append(cs, c)
	}

	return cs
}

func transformStatment(b beta.BetaStatement) CharlieStatement {

	var ids []lexeme.Token
	var expectDelim bool

	for _, tk := range b.IDs {
		switch lex := tk.Lexeme; {

		case expectDelim:
			if lex != lexeme.LEXEME_DELIM {
				panic(string("TODO: Write Err then `expected a delimiter` " + lex))
			}

		case lex == lexeme.LEXEME_ID, lex == lexeme.LEXEME_VOID:
			ids = append(ids, tk)

		default:
			panic(string("TODO: Write Err then `unexpected token` " + lex))
		}

		expectDelim = !expectDelim
	}

	return CharlieStatement{
		Assign: b.Assign,
		IDs:    ids,
		Exprs:  b.Exprs,
		Subs:   transformStatments(b.Subs),
	}
}

// PrintAll pretty prints all CharlieStatement in cs.
func PrintAll(cs []CharlieStatement) {
	printCharlieStatements(cs, 0)
	println(lexeme.LEXEME_EOF)
	println()
}

// printCharlieStatement prints all CharlieStatement in cs indenting all output
// to the specified level.
func printCharlieStatements(bs []CharlieStatement, indent int) {
	for _, b := range bs {
		printCharlieStatement(b, indent)
	}
}

// printCharlieStatement prints c indenting all output to the specified level.
func printCharlieStatement(c CharlieStatement, indent int) {

	print(strings.Repeat("  ", indent))

	if c.Assign != (lexeme.Token{}) {
		printTokens(c.IDs)
		print(" " + c.Assign.Lexeme + " ")
	}

	printTokens(c.Exprs)

	println()
	printCharlieStatements(c.Subs, indent+1)
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
