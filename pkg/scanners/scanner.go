// scanners package was created to handle scanning of tokens from a script at a
// high level; low level aspects live in the symbol package.
//
// Key decisions:
// 1. This could be rewritten to be much more performant, but I decided that
// a focus on readability was more important. Also, each script is only scanned
// once per execution so optimisation will probably not have any meaningful
// effect.
//
// This package is responsible for scanning scripts only, evaluation is
// performed by the streams/evaluator package.
package scanners

import (
	"github.com/PaulioRandall/scarlet-go/pkg/lexeme"
	"github.com/PaulioRandall/scarlet-go/pkg/scanners/matching"
)

// Method represents a scanning method
type Method string

const (
	DEFAULT          Method = `DEFAULT_SCANNER`
	PATTERN_MATCHING Method = `PATTERN_MATCHING_SCANNER`
)

// ScanAll creates a scanner from s and reads all tokens from it into an array.
func ScanAll(s string, m Method) []lexeme.Token {

	switch m {
	case DEFAULT, PATTERN_MATCHING:
		return matching.ReadAll(s)
	}

	panic(string(`Unknown scanning method '` + m + `'`))
}
