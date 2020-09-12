package compiler

import (
	"github.com/PaulioRandall/scarlet-go/inst"
	"github.com/PaulioRandall/scarlet-go/lexeme"
	"github.com/PaulioRandall/scarlet-go/number"
)

func CompileAll(con *lexeme.Container) []inst.Instruction {

	in := &input{
		in: con,
	}

	out := &output{
		out: []inst.Instruction{},
	}

	for in.more() {
		statement(in, out)
	}

	return out.out
}

func statement(in *input, out *output) {
	switch {
	case in.is(lexeme.SPELL):
		spell(in, out)

	case in.is(lexeme.ASSIGN):
		assignment(in, out)

	case in.is(lexeme.GUARD):
		guard(in, out)

	case in.is(lexeme.LOOP):
		loop(in, out)

	default:
		expression(in, out)
	}
}

func spell(in *input, out *output) {

	out.emit(inst.Instruction{
		Code:    inst.CO_DLM_PUSH,
		Snippet: in.take(),
	})

	for !in.is(lexeme.SPELL) {
		expression(in, out)
	}

	sp := in.take()
	out.emit(inst.Instruction{
		Code:    inst.CO_SPL_CALL,
		Data:    sp.Raw[1:],
		Snippet: sp,
	})
}

func assignment(in *input, out *output) {

	in.take()

	for !in.is(lexeme.ASSIGN) {
		expression(in, out)
	}

	in.take() // :=, not needed

	for first := true; first || in.is(lexeme.DELIM); first = false {

		if !first {
			in.discard() // separator
		}

		lex := in.take()

		if lex.Tok == lexeme.VOID {
			out.emit(inst.Instruction{
				Code:    inst.CO_VAL_POP,
				Snippet: lex,
			})
			return
		}

		out.emit(inst.Instruction{
			Code:    inst.CO_VAL_BIND,
			Data:    lex.Raw,
			Snippet: lex,
		})
	}
}

func guard(in *input, out *output) {

	g := in.take() // GUARD

	block := guardBlock(in)
	jumpSize := block.len()

	out.emit(inst.Instruction{
		Code:    inst.CO_JMP_FALSE,
		Data:    jumpSize,
		Snippet: g,
	})

	out.emitSet(block)
}

func guardBlock(in *input) *output {

	block := &output{
		out: []inst.Instruction{},
	}

	localBlock(in, block)
	return block
}

func loop(in *input, out *output) {

	lex := in.take() // loop

	loop := &output{
		out: []inst.Instruction{},
	}

	expression(in, loop)
	block := guardBlock(in)

	loop.emit(inst.Instruction{
		Code:    inst.CO_JMP_FALSE,
		Data:    block.len() + 1, // +1 Accounts for the jump back instruction
		Snippet: lex,
	})

	loop.emitSet(block)
	out.emitSet(loop)

	out.emit(inst.Instruction{
		Code:    inst.CO_JMP_BACK,
		Data:    loop.len() + 1, // +1 Accounts for the jump back instruction
		Snippet: lex,
	})
}

func localBlock(in *input, out *output) {

	out.emit(inst.Instruction{
		Code:    inst.CO_SUB_CTX_PUSH,
		Snippet: in.take(), // {
	})

	for !in.is(lexeme.R_CURLY) {
		statement(in, out)
	}

	out.emit(inst.Instruction{
		Code:    inst.CO_SUB_CTX_POP,
		Snippet: in.take(), // }
	})
}

func expression(in *input, out *output) {

	for {
		switch {
		case in.is(lexeme.IDENT):
			identifier(in, out)

		case in.is(lexeme.VOID):
			voidIdentifier(in, out)

		case in.tok().IsLiteral():
			literal(in, out)

		case in.tok().IsOperator():
			operator(in, out)

		case in.is(lexeme.DELIM):
			in.discard()
			return

		default:
			return
		}
	}
}

func identifier(in *input, out *output) {

	lex := in.take()

	out.emit(inst.Instruction{
		Code:    inst.CO_VAL_GET,
		Data:    lex.Raw,
		Snippet: lex,
	})
}

func voidIdentifier(in *input, out *output) {
	out.emit(inst.Instruction{
		Code:    inst.CO_VAL_PNIL,
		Snippet: in.take(),
	})
}

func literal(in *input, out *output) {

	lex := in.take()

	instruction := inst.Instruction{
		Code:    inst.CO_VAL_PVAL,
		Snippet: lex,
	}

	switch {
	case lex.Tok == lexeme.BOOL:
		instruction.Data = lex.Raw == "true"

	case lex.Tok == lexeme.NUMBER:
		instruction.Data = number.New(lex.Raw)

	case lex.Tok == lexeme.STRING:
		instruction.Data = unquote(lex.Raw)

	default:
		in.unexpected()
	}

	out.emit(instruction)
}

func operator(in *input, out *output) {

	lex := in.take()
	instruction := inst.Instruction{
		Snippet: lex,
	}

	switch {
	case lex.Tok == lexeme.ADD:
		instruction.Code = inst.CO_ADD

	case lex.Tok == lexeme.SUB:
		instruction.Code = inst.CO_SUB

	case lex.Tok == lexeme.MUL:
		instruction.Code = inst.CO_MUL

	case lex.Tok == lexeme.DIV:
		instruction.Code = inst.CO_DIV

	case lex.Tok == lexeme.REM:
		instruction.Code = inst.CO_REM

	case lex.Tok == lexeme.AND:
		instruction.Code = inst.CO_AND

	case lex.Tok == lexeme.OR:
		instruction.Code = inst.CO_OR

	case lex.Tok == lexeme.LESS:
		instruction.Code = inst.CO_LESS

	case lex.Tok == lexeme.MORE:
		instruction.Code = inst.CO_MORE

	case lex.Tok == lexeme.LESS_EQUAL:
		instruction.Code = inst.CO_LESS_EQU

	case lex.Tok == lexeme.MORE_EQUAL:
		instruction.Code = inst.CO_MORE_EQU

	case lex.Tok == lexeme.EQUAL:
		instruction.Code = inst.CO_EQU

	case lex.Tok == lexeme.NOT_EQUAL:
		instruction.Code = inst.CO_NOT_EQU

	default:
		in.unexpected()
	}

	out.emit(instruction)
}

func unquote(s string) string {

	if s == "" {
		return s
	}

	i := len(s) - 1
	return s[1:i]
}
