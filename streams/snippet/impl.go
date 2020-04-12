package snippet

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// impl is the one and only implementation of the SnippetStream interface.
type impl struct {
	ts token.TokenStream
}

// Read satisfies the SnippetStream interface.
func (uss *impl) Read() Snippet {

	tk := uss.ts.Read()
	var snip Snippet

	switch tk.Lexeme {
	case lexeme.LEXEME_EOF:
		snip.appendTokens(tk)
		snip.Kind = SNIPPET_EOF

	case lexeme.LEXEME_TERMINATOR:
		snip = uss.Read() // Consecutive terminators can be ignored

	default:
		snip.appendTokens(tk)
		uss.readNext(&snip)
	}

	return snip
}

func (uss *impl) readNext(snip *Snippet) {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR

	for tk := uss.ts.Read(); tk.Lexeme != TERMINATOR; tk = uss.ts.Read() {
		snip.appendTokens(tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			uss.readBlock(snip)
		}
	}
}

func (uss *impl) readBlock(snip *Snippet) {

	var tk lexeme.Token

	for tk = uss.ts.Read(); tk.Lexeme != lexeme.LEXEME_END; tk = uss.ts.Read() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		var sub Snippet
		uss.readNext(&sub)
		snip.appendSnippets(sub)
	}

	snip.appendTokens(tk)
}
