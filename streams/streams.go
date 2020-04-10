// streams package was created to combine multiple streams into higher level
// functionality including lexical analysis and parsing. While the individual
// parts of the scanning and parsing workflows can be executed separatly,
// primary usage is through this package's API. All functions in this package
// are pure.
//
// Key decisions: N/A
package streams

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/evaluator"
	"github.com/PaulioRandall/scarlet-go/streams/scanner"
)

// AnalyseScript performs lexical analysis (scans and evaluates) on the script
// s returning an array of tokens.
func AnalyseScript(s string) []lexeme.Token {
	tokens := scanner.ScanAll(s)
	tokens = evaluator.EvalAll(tokens)
	return tokens
}
