package compiler

import (
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/code"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/inst"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/eskarina/prop"
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

	defer com.reject() // GEN_TERMINATOR, now redundant

	switch {
	case com.empty():
		com.unexpected()

	case com.has(prop.PR_CALLABLE):
		call(com)

	default:
		com.unexpected()
	}
}

func call(com *compiler) {

	com.take() // PR_PARAMETERS redundant
	argCount := 0

	for !com.has(prop.PR_SPELL) {
		argCount++
		expression(com)
	}

	sp := com.take()

	com.output(&inst.Instruction{
		Code:    code.CO_VAL_PUSH,
		Data:    argCount,
		Snippet: sp,
	})

	com.output(&inst.Instruction{
		Code:    code.CO_SPELL,
		Data:    sp.Raw[1:],
		Snippet: sp,
	})
}

func expression(com *compiler) {

	switch {
	case com.has(prop.PR_IDENTIFIER):
		identifier(com)

	case com.has(prop.PR_LITERAL):
		literal(com)

	default:
		com.unexpected()
	}
}

func identifier(com *compiler) {

	lex := com.take()

	com.output(&inst.Instruction{
		Code:    code.CO_CTX_GET,
		Data:    lex.Raw,
		Snippet: lex,
	})
}

func literal(com *compiler) {

	lex := com.take()

	in := &inst.Instruction{
		Code:    code.CO_VAL_PUSH,
		Snippet: lex,
	}

	switch {
	case lex.Is(prop.PR_BOOL):
		in.Data = lex.Raw == "true"

	case lex.Is(prop.PR_NUMBER):
		//in.Data = number.New(lex.Raw())

	case lex.Is(prop.PR_STRING):
		in.Data = lex.Raw

	default:
		com.unexpected()
	}

	com.output(in)
}
