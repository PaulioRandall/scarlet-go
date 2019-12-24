package scanner

import (
	"strings"

	"github.com/PaulioRandall/scarlet-go/token"
)

// findSymbol satisfies the source.TokenFinder function prototype. It attempts
// to match the next token to a symbol kind returning its length if matched.
func findSymbol(r []rune) (_ int, _ token.Kind, _ error) {

	type sym struct {
		v string
		n int
		k token.Kind
	}

	symbols := []sym{
		sym{":=", 2, token.ASSIGN},
		sym{"->", 2, token.RETURNS},
		sym{"(", 1, token.OPEN_PAREN},
		sym{")", 1, token.CLOSE_PAREN},
		sym{",", 1, token.ID_DELIM},
		sym{"@", 1, token.SPELL},
		sym{"{", 1, token.OPEN_LIST},
		sym{"}", 1, token.CLOSE_LIST},
		sym{"+", 1, token.ADD},
		sym{"-", 1, token.SUBTRACT},
		sym{"/", 1, token.DIVIDE},
		sym{"*", 1, token.MULTIPLY},
		sym{"%", 1, token.MODULO},
		sym{"|", 1, token.OR},
		sym{"&", 1, token.AND},
		sym{"=", 1, token.EQUAL},
		sym{"#", 1, token.NOT_EQUAL},
		sym{"<=", 2, token.LT_OR_EQUAL},
		sym{">=", 2, token.GT_OR_EQUAL},
		sym{"<", 1, token.LT},
		sym{">", 1, token.GT},
	}

	src := string(r)

	for _, s := range symbols {
		if strings.HasPrefix(src, s.v) {
			return s.n, s.k, nil
		}
	}

	return
}
