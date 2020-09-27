package temp

import (
	old "github.com/PaulioRandall/scarlet-go/lexeme"
	new "github.com/PaulioRandall/scarlet-go/token/container"
	"github.com/PaulioRandall/scarlet-go/token/lexeme"
)

// ConvertContainer creates an old style lexeme container from the new style.
func ConvertContainer(in *new.Container) *old.Container {

	out := old.NewContainer(nil)
	for it := in.Iterator(); it.HasNext(); {
		lexNew := it.Next()
		lexOld := ConvertLexeme(lexNew)
		out.Put(lexOld)
	}

	return out
}

// ConvertLexeme creates an old style lexeme from the new style.
func ConvertLexeme(l lexeme.Lexeme) *old.Lexeme {
	return &old.Lexeme{
		Tok:  tokenMapping[l.Type()],
		Raw:  l.Raw(),
		Line: l.Line(),
		Col:  l.Col(),
	}
}

var tokenMapping = map[lexeme.TokenType]old.Token{
	lexeme.SPACE:      old.SPACE,
	lexeme.COMMENT:    old.COMMENT,
	lexeme.TERMINATOR: old.TERMINATOR,
	lexeme.NEWLINE:    old.NEWLINE,
	lexeme.BOOL:       old.BOOL,
	lexeme.NUMBER:     old.NUMBER,
	lexeme.STRING:     old.STRING,
	lexeme.IDENT:      old.IDENT,
	lexeme.SPELL:      old.SPELL,
	lexeme.GUARD:      old.GUARD,
	lexeme.LOOP:       old.LOOP,
	lexeme.DELIM:      old.DELIM,
	lexeme.L_PAREN:    old.L_PAREN,
	lexeme.R_PAREN:    old.R_PAREN,
	lexeme.L_SQUARE:   old.L_SQUARE,
	lexeme.R_SQUARE:   old.R_SQUARE,
	lexeme.L_CURLY:    old.L_CURLY,
	lexeme.R_CURLY:    old.R_CURLY,
	lexeme.ASSIGN:     old.ASSIGN,
	lexeme.VOID:       old.VOID,
	lexeme.ADD:        old.ADD,
	lexeme.SUB:        old.SUB,
	lexeme.MUL:        old.MUL,
	lexeme.DIV:        old.DIV,
	lexeme.REM:        old.REM,
	lexeme.AND:        old.AND,
	lexeme.OR:         old.OR,
	lexeme.LESS:       old.LESS,
	lexeme.MORE:       old.MORE,
	lexeme.LESS_EQUAL: old.LESS_EQUAL,
	lexeme.MORE_EQUAL: old.MORE_EQUAL,
	lexeme.EQUAL:      old.EQUAL,
	lexeme.NOT_EQUAL:  old.NOT_EQUAL,
}
