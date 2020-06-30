package group

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type GroupFunc func() ([]Token, GroupFunc, error)

type TokenStream interface {
	Next() Token
}

type Group struct {
}
