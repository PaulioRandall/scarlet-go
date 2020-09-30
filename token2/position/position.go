package position

import (
	"fmt"
)

type Position struct {
	Offset  int // Bytes
	Line    int // Index
	ColByte int // Index
	ColRune int // Index
}

func (p Position) String() string {
	return fmt.Sprintf("[%d]%d/%d/%d",
		p.Line,
		p.Offset,
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
