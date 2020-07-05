package compile

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type TokenStream interface {
	Next() Token
}

func compile(ts TokenStream) ([]instruction, error) {
	return nil, nil
}
