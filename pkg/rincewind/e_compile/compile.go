package compile

/*
import (
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/inst"
	. "github.com/PaulioRandall/scarlet-go/pkg/rincewind/token"
)

type CompileFunc func() (Instruction, CompileFunc, error)

type TokenStream interface {
	Next() Token
}

type pipeAndMatch struct {
	ts TokenStream
}

type compiler struct {
	*PipeStack
}

func New(ts TokenStream) CompileFunc {

	if ts == nil {
		failNow("Non-nil TokenStream required")
	}

	pam := pipeAndMatch{ts}
	com := &compiler{
		NewPipeStack(pam, pam),
	}

	if com.empty() {
		return nil
	}

	return com.compile
}

func (com *compiler) compile() (Instruction, CompileFunc, error) {

	in, e := next(com)
	if e != nil {
		return nil, nil, e
	}

	if rfx.Empty() {
		return tk, nil, nil
	}

	return tk, rfx.refix, nil
}

func (com *compiler) next() Token {

	if com.empty() {
		failNow("No tokens remaining, call `match` or `empty` first")
	}

	r := com.buff
	com.buff = com.ts.Next()

	return r
}


func (com *compiler) empty() bool {
	return com.buff == nil
}

func (com *compiler) match(gen GenType, sub SubType) bool {

	if com.empty() {
		return false
	}

	g := gen == GE_ANY || gen == com.buff.GenType()
	s := sub == SU_ANY || sub == com.buff.SubType()

	return g && s
}

func (com *compiler) notMatch(gen GenType, sub SubType) bool {
	return !com.match(gen, sub)
}

func (com *compiler) expect(gen GenType, sub SubType) (Token, error) {

	if com.empty() {
		return nil, errorUnexpectedEOF(com)
	}

	if com.match(gen, sub) {
		return com.next(), nil
	}

	return nil, errorWrongToken(com, com.buff)
}
*/
