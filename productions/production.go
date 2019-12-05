package productions

import (
	"github.com/PaulioRandall/scarlet-go/perror"
	//"github.com/PaulioRandall/scarlet-go/where"
	"github.com/PaulioRandall/scarlet-go/source"
	"github.com/PaulioRandall/scarlet-go/token"
)

// Production represents a production rule.
type Production interface {
	Next() (token.Token, Production, perror.Perror)
}

type rule struct {
	parent *Production
	src    *source.Source
}

// newline is a terminal production rule that will always return a newline token
// and its parent production.
type newline string
