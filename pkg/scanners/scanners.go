package scanners

import (
	"github.com/PaulioRandall/scarlet-go/pkg/scanners/matching"
	"github.com/PaulioRandall/scarlet-go/pkg/token"
)

type Method string

const (
	DEFAULT          Method = `DEFAULT_SCANNER`
	PATTERN_MATCHING Method = `PATTERN_MATCHING_SCANNER`
)

func ScanAll(s string, m Method) []token.Token {

	switch m {
	case DEFAULT, PATTERN_MATCHING:
		return matching.ScanAll(s)
	}

	panic(string(`Unknown scanning method '` + m + `'`))
}
