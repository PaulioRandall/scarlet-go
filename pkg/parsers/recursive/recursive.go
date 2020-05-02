package recursive

import (
	"github.com/PaulioRandall/scarlet-go/pkg/token"

	st "github.com/PaulioRandall/scarlet-go/pkg/statement"
)

func ParseAll(tks []token.Token) []st.Statement {
	p := &pipe{token.NewIterator(tks)}
	return parseStatements(p)
}
