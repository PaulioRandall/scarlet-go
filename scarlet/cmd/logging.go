package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

func makeLogFilename(c config, ext string) string {
	ext = "." + ext
	f := filepath.Base(c.script)
	f = strings.TrimSuffix(f, filepath.Ext(f))
	return filepath.Join(c.logDir, f+ext)
}

func logContainer(c config, con *lexeme.Container, ext string) error {
	filename := makeLogFilename(c, ext)
	return writeLexemeFile(filename, con.Head())
}

func writeLexemeFile(filename string, head *lexeme.Lexeme) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return lexeme.Print(f, head)
}

func logInstructions(c config, ins []inst.Instruction, ext string) error {
	filename := makeLogFilename(c, ext)
	return writeInstructionFile(filename, ins)
}

func writeInstructionFile(filename string, ins []inst.Instruction) error {

	f, e := os.Create(filename)
	if e != nil {
		return e
	}

	defer f.Close()
	return inst.Print(f, ins)
}
