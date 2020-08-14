package compiler

import (
	"github.com/PaulioRandall/scarlet-go/shared/inst"
	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/shared/number"
)

func CompileAll(con *lexeme.Container) []inst.Instruction {

	com := &compiler{
		input: con,
		out:   []inst.Instruction{},
	}

	compile(com)
	return com.out
}

func compile(com *compiler) {

	for com.more() {

		switch {
		case com.empty():
			com.unexpected()

		case com.is(lexeme.SPELL):
			spell(com)

		case com.is(lexeme.ASSIGN):
			assignment(com)

		default:
			expression(com)
		}

		com.reject() // GEN_TERMINATOR, now redundant
	}
}

func spell(com *compiler) {

	com.output(inst.Instruction{
		Code:    inst.CO_DELIM_PUSH,
		Snippet: com.take(),
	})

	for !com.is(lexeme.SPELL) {
		expression(com)
	}

	sp := com.take()
	com.output(inst.Instruction{
		Code:    inst.CO_SPELL,
		Data:    sp.Raw[1:],
		Snippet: sp,
	})
}

func assignment(com *compiler) {

	com.take()

	for !com.is(lexeme.ASSIGN) {
		expression(com)
	}

	com.take() // :=, not needed

	for first := true; first || com.is(lexeme.DELIM); first = false {

		if !first {
			com.reject() // separator
		}

		lex := com.take()

		com.output(inst.Instruction{
			Code:    inst.CO_CTX_SET,
			Data:    lex.Raw,
			Snippet: lex,
		})
	}
}

func expression(com *compiler) {

	for {
		switch {
		case com.is(lexeme.IDENT):
			identifier(com)

		case com.tok().IsLiteral():
			literal(com)

		case com.tok().IsOperator():
			operator(com)

		case com.is(lexeme.DELIM):
			com.reject()
			return

		default:
			return
		}
	}
}

func identifier(com *compiler) {

	lex := com.take()

	com.output(inst.Instruction{
		Code:    inst.CO_CTX_GET,
		Data:    lex.Raw,
		Snippet: lex,
	})
}

func literal(com *compiler) {

	lex := com.take()

	in := inst.Instruction{
		Code:    inst.CO_VAL_PUSH,
		Snippet: lex,
	}

	switch {
	case lex.Tok == lexeme.BOOL:
		in.Data = lex.Raw == "true"

	case lex.Tok == lexeme.NUMBER:
		in.Data = number.New(lex.Raw)

	case lex.Tok == lexeme.STRING:
		in.Data = unquote(lex.Raw)

	default:
		com.unexpected()
	}

	com.output(in)
}

func operator(com *compiler) {

	lex := com.take()
	in := inst.Instruction{
		Snippet: lex,
	}

	switch {
	case lex.Tok == lexeme.ADD:
		in.Code = inst.CO_ADD

	case lex.Tok == lexeme.SUB:
		in.Code = inst.CO_SUB

	case lex.Tok == lexeme.MUL:
		in.Code = inst.CO_MUL

	case lex.Tok == lexeme.DIV:
		in.Code = inst.CO_DIV

	case lex.Tok == lexeme.REM:
		in.Code = inst.CO_REM

	case lex.Tok == lexeme.AND:
		in.Code = inst.CO_AND

	case lex.Tok == lexeme.OR:
		in.Code = inst.CO_OR

	case lex.Tok == lexeme.LESS:
		in.Code = inst.CO_LESS

	case lex.Tok == lexeme.MORE:
		in.Code = inst.CO_MORE

	case lex.Tok == lexeme.LESS_EQUAL:
		in.Code = inst.CO_LESS_EQU

	case lex.Tok == lexeme.MORE_EQUAL:
		in.Code = inst.CO_MORE_EQU

	default:
		com.unexpected()
	}

	com.output(in)
}

func unquote(s string) string {

	if s == "" {
		return s
	}

	i := len(s) - 1
	return s[1:i]
}
