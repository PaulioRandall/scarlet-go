package compiler

import (
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/inst"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
	"github.com/PaulioRandall/scarlet-go/eskarina/shared/number"
)

func CompileAll(head *lexeme.Lexeme) *inst.Instruction {

	com := &compiler{
		input: lexeme.NewContainer(head),
		out:   &inst.Container{},
	}

	compile(com)
	return com.out.Head()
}

func compile(com *compiler) {

	for com.more() {

		switch {
		case com.empty():
			com.unexpected()

		case com.has(lexeme.PR_CALLABLE):
			call(com)

		default:
			com.unexpected()
		}

		com.reject() // GEN_TERMINATOR, now redundant
	}
}

func call(com *compiler) {

	com.take() // PR_PARAMETERS redundant
	argCount := 0

	for !com.has(lexeme.PR_SPELL) {
		argCount++
		expression(com)
	}

	sp := com.take()

	com.output(&inst.Instruction{
		Code:    inst.CO_VAL_PUSH,
		Data:    argCount,
		Snippet: sp,
	})

	com.output(&inst.Instruction{
		Code:    inst.CO_SPELL,
		Data:    sp.Raw[1:],
		Snippet: sp,
	})
}

func expression(com *compiler) {

	switch {
	case com.has(lexeme.PR_IDENTIFIER):
		identifier(com)

	case com.has(lexeme.PR_LITERAL):
		literal(com)

	default:
		com.unexpected()
	}
}

func identifier(com *compiler) {

	lex := com.take()

	com.output(&inst.Instruction{
		Code:    inst.CO_CTX_GET,
		Data:    lex.Raw,
		Snippet: lex,
	})
}

func literal(com *compiler) {

	lex := com.take()

	in := &inst.Instruction{
		Code:    inst.CO_VAL_PUSH,
		Snippet: lex,
	}

	switch {
	case lex.Is(lexeme.PR_BOOL):
		in.Data = lex.Raw == "true"

	case lex.Is(lexeme.PR_NUMBER):
		in.Data = number.New(lex.Raw)

	case lex.Is(lexeme.PR_STRING):
		in.Data = unquote(lex.Raw)

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