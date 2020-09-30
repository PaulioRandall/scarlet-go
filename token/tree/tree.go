package tree

import (
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

type Node interface {
}

type Tree struct {
}

type IdentNode struct {
	tk lexeme.Lexeme
}
