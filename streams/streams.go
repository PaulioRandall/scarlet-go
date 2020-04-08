// streams package provides a set of components that are to be piped together to
// form a Scarlet interpreter.
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
