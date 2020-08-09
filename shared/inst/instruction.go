package inst

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/shared/lexeme"
)

type Instruction struct {
	Code    Code
	Data    interface{}
	Snippet *lexeme.Lexeme
}

func (in Instruction) String() string {
	return fmt.Sprintf("%v %v",
		in.Code.String(),
		in.Data,
	)
}
