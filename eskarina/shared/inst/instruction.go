package inst

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/eskarina/shared/lexeme"
)

type Instruction struct {
	Code    Code
	Data    interface{}
	Snippet *lexeme.Lexeme
	Next    *Instruction
}

func (in Instruction) NextNode() *Instruction {
	return in.Next
}

func (in Instruction) String() string {

	lineBegin, lineEnd, colBegin, colEnd := in.snippet()

	return fmt.Sprintf("%d:%d->%d:%d %v %v",
		lineBegin+1, colBegin,
		lineEnd+1, colEnd,
		in.Code.String(),
		in.Data,
	)
}

func (in Instruction) snippet() (lineBegin, colBegin, lineEnd, colEnd int) {

	const max_int = int(^uint(0) >> 1)

	lineBegin = max_int
	lineEnd = 0
	colBegin = max_int
	colEnd = 0

	it := lexeme.NewItinerant(in.Snippet).To().Iterator()
	for it.Next() {
		lex := it.Curr()

		switch {
		case lex.Line < lineBegin:
			lineBegin = lex.Line
			colBegin = lex.Col

		case lex.Line == lineBegin && lex.Col < colBegin:
			colBegin = lex.Col
		}

		switch {
		case lex.Line > lineBegin:
			lineEnd = lex.Line
			colEnd = lex.Col

		case lex.Line == lineEnd && lex.Col > colEnd:
			colEnd = lex.Col
		}
	}

	return
}

func (in *Instruction) ToSlice() []Instruction {

	var ins []Instruction

	for next := in; next != nil; next = next.Next {
		ins = append(ins, *next)
	}

	return ins
}
