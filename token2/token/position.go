package token

import (
	"fmt"
)

type Position struct {
	SrcOffset int // Bytes
	LineIdx   int
	ColByte   int
	ColRune   int
}

func (p Position) String() string {
	return fmt.Sprintf("[%d]%d/%d/%d",
		p.LineIdx,
		p.SrcOffset,
		p.ColByte,
		p.ColRune,
	)
}

type Snippet struct {
	Position
	End Position
}

func (s Snippet) String() string {
	return fmt.Sprintf("%s -> %s", s.Position.String(), s.End.String())
}
