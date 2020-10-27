package inst

import (
	"fmt"

	"github.com/PaulioRandall/scarlet-go/token/code"
	"github.com/PaulioRandall/scarlet-go/token/value"
)

// Inst represents an instruction, data may be included with some codes.
type Inst struct {
	Code code.Code
	Data value.Value
}

// String returns a human readable string representation of the instruction.
func (in Inst) String() string {
	if in.Data == nil {
		return in.Code.String()
	}
	return fmt.Sprintf("%s %s", in.Code.String(), in.Data.String())
}
