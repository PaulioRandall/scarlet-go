package parser

import (
	"github.com/PaulioRandall/scarlet-go/lexeme"
)

// Parser is parser for a stream of tokens.
type Parser struct {
	in  chan lexeme.Token
	buf *lexeme.Token // buffer, allows for look ahead by 1
}

// New creates a new Parser.
func New(in chan lexeme.Token) *Parser {
	return &Parser{
		in: in,
	}
}

// Parse parses tokens obtained via the input channel into statements until the
// channel is closed. A master statement is returned that represents the block
// of statements.
func (p *Parser) Parse() (_ Stat) {
	return p.parseStats(lexeme.Token{
		Lexeme: lexeme.LEXEME_SOF,
	})
}

// peek returns the token on the buffer. If the buffer is empty the the next
// token is read from the input channel into the buffer first. Panics if the
// channel is closed when attempting a read.
func (p *Parser) peek() lexeme.Token {

	if p.buf == nil {
		tk, ok := <-p.in

		if !ok {
			panic(newTkErr(tk, "Token input channel closed prematurely"))
		}

		p.buf = &tk
	}

	return *p.buf
}

// take removes and returns the token in the buffer. If the buffer is empty then
// a peek operation is performed to fill it first.
func (p *Parser) take() (tk lexeme.Token) {

	if p.buf == nil {
		p.peek()
	}

	tk, p.buf = *p.buf, nil
	return
}

// ensure will panic if the specified token is not one of the specified kinds.
func (p *Parser) ensure(tk lexeme.Token, lexs ...lexeme.Lexeme) {

	var errMsg string
	tkLex := tk.Lexeme

	for _, lex := range lexs {
		if tkLex == lex {
			return
		}
	}

	errMsg = "Expected "
	if len(lexs) == 1 {
		errMsg += string(lexs[0])
	} else {
		errMsg = "either"
		for _, lex := range lexs {
			errMsg += " " + string(lex)
		}
	}

	errMsg += " but was " + string(tk.Lexeme)
	panic(newTkErr(tk, errMsg))
}

// peekEnsure returns the next token in the input channel but will panic if
// the if the channel is closed or the specified token is not one of the
// specified kinds.
func (p *Parser) peekEnsure(lexs ...lexeme.Lexeme) lexeme.Token {
	tk := p.peek()
	p.ensure(tk, lexs...)
	return tk
}

// takeEnsure returns the next token in the input channel but will panic if
// the if the channel is closed or the specified token is not one of the
// specified kinds.
func (p *Parser) takeEnsure(lexs ...lexeme.Lexeme) lexeme.Token {
	tk := p.take()
	p.ensure(tk, lexs...)
	return tk
}
