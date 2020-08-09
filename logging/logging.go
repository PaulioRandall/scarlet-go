package logging

import (
	"os"

	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

func writeLexemeFile(filename string, head *lexeme.Lexeme) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return lexeme.Print(f, head)
}

func writeInstructionFile(filename string, ins []inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.Print(f, ins)
}
