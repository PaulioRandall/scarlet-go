package sanitiser

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	//"github.com/PaulioRandall/scarlet-go/pkg/eskarina/perror"
	//"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
)

func SanitiseAll(first *lexeme.Lexeme) (*lexeme.Lexeme, error) {

	for lex := first; lex != nil; lex = lex.Next {

		keep, e := sanitise(lex)
		if e != nil {
			return nil, e
		}

		if !keep {

		}
	}

	return nil, nil
}

func sanitise(lex *lexeme.Lexeme) (bool, error) {
	return true, nil
}
