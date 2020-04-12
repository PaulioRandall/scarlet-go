package snippet

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// impl is the one and only implementation of the SnippetStream interface.
type impl struct {
	ts  token.TokenStream
	buf *lexeme.Token
}

// Read satisfies the SnippetStream interface.
func (uss *impl) Read() Snippet {

	tk := uss.peekToken()
	var snip Snippet

	switch tk.Lexeme {
	case lexeme.LEXEME_EOF:
		snip.Kind = SNIPPET_EOF

	case lexeme.LEXEME_TERMINATOR:
		snip = uss.Read() // Consecutive terminators can be ignored

	default:
		uss.readStatement(&snip)
	}

	return snip
}

func (uss *impl) readToken() lexeme.Token {
	tk := uss.peekToken()
	uss.buf = nil
	return tk
}

func (uss *impl) peekToken() lexeme.Token {

	if uss.buf == nil {
		tk := uss.ts.Read()
		uss.buf = &tk
		return tk
	}

	tk := *(uss.buf)
	return tk
}

func (uss *impl) readStatement(snip *Snippet) {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR

	for tk := uss.readToken(); tk.Lexeme != TERMINATOR; tk = uss.readToken() {
		snip.appendTokens(tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			uss.readBlock(snip)
		}
	}
}

func (uss *impl) readBlock(snip *Snippet) {

	var tk lexeme.Token

	for tk = uss.peekToken(); tk.Lexeme != lexeme.LEXEME_END; tk = uss.peekToken() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		var sub Snippet
		uss.readStatement(&sub)
		snip.appendSnippets(sub)
	}

	snip.appendTokens(tk)
}
