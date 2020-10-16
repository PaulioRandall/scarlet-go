package inst

import (
	"fmt"
)

type (
	// DataRef represents a reference to a data value.
	DataRef uint16

	// RiscInst represents a RISC instruction with a reference to any accompanying
	// data.
	RiscInst struct {
		Inst Inst
		Data DataRef
	}
)

const NoData DataRef = 0

// String returns a human readable string representation of the instruction.
func (ri RiscInst) String() string {
	return fmt.Sprintf("%d %d", ri.Inst, ri.Data)
}
