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

		case com.is(lexeme.CALLABLE):
			call(com)

		case com.is(lexeme.ASSIGNMENT):
			assignment(com)

		default:
			expression(com)
		}

		com.reject() // GEN_TERMINATOR, now redundant
	}
}

func call(com *compiler) {

	com.take() // Now redundant
	argCount := 0

	for !com.is(lexeme.SPELL) {
		argCount++
		expression(com)
	}

	sp := com.take()

	com.output(inst.Instruction{
		Code:    inst.CO_VAL_PUSH,
		Data:    argCount,
		Snippet: sp,
	})

	com.output(inst.Instruction{
		Code:    inst.CO_SPELL,
		Data:    sp.Raw[1:],
		Snippet: sp,
	})
}

func assignment(com *compiler) {

	com.take() // Now redundant

	for !com.is(lexeme.ASSIGNMENT) {
		expression(com)
	}

	com.take() // :=, not needed

	for first := true; first || com.is(lexeme.SEPARATOR); first = false {

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
		case com.is(lexeme.IDENTIFIER):
			identifier(com)

		case com.tok().IsLiteral():
			literal(com)

		case com.tok().IsOperator():
			operator(com)

		case com.is(lexeme.SEPARATOR):
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
	// TODO
	com.take()
}

func unquote(s string) string {

	if s == "" {
		return s
	}

	i := len(s) - 1
	return s[1:i]
}
