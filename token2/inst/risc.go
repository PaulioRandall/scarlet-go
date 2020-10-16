package inst

import (
	"fmt"
)

type DataRef uint16

// RiscInst represents a RISC instruction with a reference to any accompanying
// data.
type RiscInst struct {
	Inst Inst
	Data DataRef
}

// HasData returns true if the instruction references some data in the static
// data pool.
func (ri RiscInst) HasData() bool {
	return ri.Data > 0
}

// String returns a human readable string representation of the instruction.
func (ri RiscInst) String() string {
	return fmt.Sprintf("%d %d", ri.Inst, ri.Data)
}
