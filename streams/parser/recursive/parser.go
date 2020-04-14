package recursive

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

func ParseAll(tks []lexeme.Token) Tree {
	itr := lexeme.NewIterator(tks)
	return Tree(parseStatements(itr))
}

func parseStatements(itr *lexeme.TokenIterator) []Statement {
	return []Statement{}
}

func parseStatement(itr *lexeme.TokenIterator) Statement {
	return Statement{}
}
