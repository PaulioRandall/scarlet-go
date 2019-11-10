package main

import (
	"github.com/PaulioRandall/scarlet-go/scanner"
)

func main() {
	s := scanner.New("  \n       \r\n")
	tok, s := s()

	for tok != scanner.EmptyTok() {
		print(tok.Kind.Name() + " ")
		tok, s = s()
	}
}
