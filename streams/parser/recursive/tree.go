package recursive

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

type Tree []Statement

type Statement struct {
	IDs    []lexeme.Token
	Assign lexeme.Token
	Exprs  []Expression
}

type Kind string

const (
	EXPR_ARITHMETIC Kind = `ARITHMETIC`
	EXPR_LOGIC      Kind = `LOGIC`
	EXPR_FUNC_CALL  Kind = `FUNC_CALL`
	EXPR_SPELL_CALL Kind = `SPELL_CALL`
)

type Expression interface {
	Kind() Kind
}

type Arithmetic struct {
	Left     Expression
	Operator lexeme.Token
	Right    Expression
}

func (_ *Arithmetic) Kind() Kind {
	return EXPR_ARITHMETIC
}

type Logic struct {
	Left     Expression
	Operator lexeme.Token
	Right    Expression
}

func (_ *Logic) Kind() Kind {
	return EXPR_LOGIC
}

type FuncCall struct {
	ID     lexeme.Token
	Input  []lexeme.Token
	Output []lexeme.Token
}

func (_ *FuncCall) Kind() Kind {
	return EXPR_FUNC_CALL
}

type SpellCall struct {
	Spell  lexeme.Token
	ID     lexeme.Token
	Input  []lexeme.Token
	Output []lexeme.Token
}

func (_ *SpellCall) Kind() Kind {
	return EXPR_SPELL_CALL
}
