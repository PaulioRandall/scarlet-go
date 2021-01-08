package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTextMarker_Advance(t *testing.T) {

	var (
		tm      = TextMarker{}
		offset  = 0
		eng     = "I don't understand"
		engByte = len(eng)         // 18
		engRune = len([]rune(eng)) // 18
	)

	tm.Advance(eng)
	offset += engByte
	require.Equal(t, offset, tm.Offset)
	require.Equal(t, engByte, tm.ColByte)
	require.Equal(t, engRune, tm.ColRune)

	tm.Advance("\n")
	offset++
	require.Equal(t, offset, tm.Offset)
	require.Equal(t, 1, tm.Line)
	require.Equal(t, 0, tm.ColByte)
	require.Equal(t, 0, tm.ColRune)

	var (
		jap     = "日本語"
		japByte = len(jap)         // 9
		japRune = len([]rune(jap)) // 3
	)

	tm.Advance(jap)
	offset += japByte
	require.Equal(t, offset, tm.Offset)
	require.Equal(t, 1, tm.Line)
	require.Equal(t, japByte, tm.ColByte)
	require.Equal(t, japRune, tm.ColRune)

	var (
		abc          = "\nabc\nefg"
		abcByte      = len(abc) // 7
		bytesAfterLF = 3
		runesAfterLF = 3
	)

	tm.Advance(abc)
	offset += abcByte
	require.Equal(t, offset, tm.Offset)
	require.Equal(t, 3, tm.Line)
	require.Equal(t, bytesAfterLF, tm.ColByte)
	require.Equal(t, runesAfterLF, tm.ColRune)
}
