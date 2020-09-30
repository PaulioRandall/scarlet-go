package token

type Position struct {
	File    string
	Offset  int
	LineIdx int
	ByteCol int
	RuneCol int
}
