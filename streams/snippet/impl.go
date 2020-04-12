package snippet

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// impl is the one and only implementation of the SnippetStream interface.
type impl struct {
	ts  token.TokenStream // Do not access directly, use readToken & peekToken
	buf *lexeme.Token
}

// Read satisfies the SnippetStream interface.
func (uss *impl) Read() Snippet {

	lex := uss.peekToken().Lexeme

	if lex == lexeme.LEXEME_EOF {
		return Snippet{}
	}

	if lex == lexeme.LEXEME_TERMINATOR {
		return uss.Read() // Consecutive terminators can be ignored
	}

	return uss.readSnippet()
}

// readToken reads a token from the underlying TokenStream. The TokenStream
// must NOT be accessed directly as buffer management is required.
func (uss *impl) readToken() lexeme.Token {
	tk := uss.peekToken()
	uss.buf = nil
	return tk
}

// peekToken reads a token from the underlying TokenStream. The token is cached
// so further calls to peekToken or a subsequent call to readToken will return
// the peeked token rather than a new one. The TokenStream must NOT be accessed
// directly as buffer management is required.
func (uss *impl) peekToken() lexeme.Token {

	if uss.buf == nil {
		tk := uss.ts.Read()
		uss.buf = &tk
		return tk
	}

	tk := *(uss.buf)
	return tk
}

// readSnippet reads all tokens representing the next statement from the
// TokenStream and returns a new snippet.
func (uss *impl) readSnippet() Snippet {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR
	var snip Snippet

	for tk := uss.readToken(); tk.Lexeme != TERMINATOR; tk = uss.readToken() {
		snip.Tokens = append(snip.Tokens, tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			snip.Snippets = uss.readBlock()
		}
	}

	return snip
}

// readBlock reads all statements in the current block into an array.
func (uss *impl) readBlock() []Snippet {

	var snips []Snippet

	for tk := uss.peekToken(); tk.Lexeme != lexeme.LEXEME_END; tk = uss.peekToken() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		sub := uss.readSnippet()
		snips = append(snips, sub)
	}

	return snips
}
