package statement

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"

	"github.com/PaulioRandall/scarlet-go/streams/token"
)

// impl is the one and only implementation of the StatementStream interface.
type impl struct {
	ts  token.TokenStream // Do not access directly, use readToken & peekToken
	buf *lexeme.Token
}

// Read satisfies the StatementStream interface.
func (uss *impl) Read() Statement {

	lex := uss.peekToken().Lexeme

	if lex == lexeme.LEXEME_EOF {
		return Statement{}
	}

	if lex == lexeme.LEXEME_TERMINATOR {
		return uss.Read() // Consecutive terminators can be ignored
	}

	return uss.readStatement()
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

// readStatement reads all tokens representing the next statement from the
// TokenStream and returns a new Statement.
func (uss *impl) readStatement() Statement {

	const TERMINATOR = lexeme.LEXEME_TERMINATOR
	var stat Statement

	for tk := uss.readToken(); tk.Lexeme != TERMINATOR; tk = uss.readToken() {
		stat.Tokens = append(stat.Tokens, tk)

		if tk.Lexeme == lexeme.LEXEME_DO {
			stat.Stats = uss.readBlock()
		}
	}

	return stat
}

// readBlock reads all statements in the current block into an array.
func (uss *impl) readBlock() []Statement {

	var subs []Statement

	for tk := uss.peekToken(); tk.Lexeme != lexeme.LEXEME_END; tk = uss.peekToken() {

		if tk.Lexeme == lexeme.LEXEME_TERMINATOR || tk.Lexeme == lexeme.LEXEME_EOF {
			panic(newErr(tk, `Expected a statement or 'END', not '%s'`, tk.Value))
		}

		sub := uss.readStatement()
		subs = append(subs, sub)
	}

	return subs
}
