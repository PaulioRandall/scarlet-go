package lexeme

import (
	"unicode"
)

// Lexeme represents a token type.
type Lexeme string

const (
	LEXEME_UNDEFINED Lexeme = ``
	// ------------------
	LEXEME_SOF         Lexeme = `SOF`
	LEXEME_EOF         Lexeme = `EOF`
	LEXEME_COMMENT     Lexeme = `COMMENT`
	LEXEME_WHITESPACE  Lexeme = `WHITESPACE`
	LEXEME_NEWLINE     Lexeme = `NEWLINE`
	LEXEME_FUNC        Lexeme = `FUNCTION_KEYWORD`
	LEXEME_MATCH       Lexeme = `MATCH_BLOCK_KEYWORD`
	LEXEME_INLINE      Lexeme = `INLINE_BLOCK_IMPLIED`
	LEXEME_ID          Lexeme = `ID`
	LEXEME_DELIM       Lexeme = `DELIMITER`
	LEXEME_ASSIGN      Lexeme = `ASSIGNMENT`
	LEXEME_RETURNS     Lexeme = `RETURNS`
	LEXEME_DO          Lexeme = `BLOCK_OPEN`
	LEXEME_END         Lexeme = `BLOCK_CLOSE`
	LEXEME_OPEN_PAREN  Lexeme = `PAREN_OPEN`
	LEXEME_CLOSE_PAREN Lexeme = `PAREN_CLOSE`
	LEXEME_OPEN_GUARD  Lexeme = `GUARD_OPEN`
	LEXEME_CLOSE_GUARD Lexeme = `GUARD_CLOSE`
	LEXEME_OPEN_LIST   Lexeme = `LIST_OPEN`
	LEXEME_CLOSE_LIST  Lexeme = `LIST_CLOSE`
	LEXEME_SPELL       Lexeme = `SPELL`
	LEXEME_STRING      Lexeme = `STRING`
	LEXEME_TEMPLATE    Lexeme = `TEMPLATE`
	LEXEME_INT         Lexeme = `INT`
	LEXEME_FLOAT       Lexeme = `FLOAT`
	LEXEME_BOOL        Lexeme = `BOOL`
	LEXEME_NOT         Lexeme = `NOT`
	LEXEME_ADD         Lexeme = `ADD`
	LEXEME_SUBTRACT    Lexeme = `SUBTRACT`
	LEXEME_MULTIPLY    Lexeme = `MULTIPLY`
	LEXEME_DIVIDE      Lexeme = `DIVIDE`
	LEXEME_REMAINDER   Lexeme = `REMAINDER`
	LEXEME_AND         Lexeme = `AND`
	LEXEME_OR          Lexeme = `OR`
	LEXEME_EQU         Lexeme = `EQUAL`
	LEXEME_NEQ         Lexeme = `NOT_EQUAL`
	LEXEME_LT          Lexeme = `LESS_THAN`
	LEXEME_LT_OR_EQU   Lexeme = `LESS_THAN_OR_EQUAL`
	LEXEME_MT          Lexeme = `MORE_THAN`
	LEXEME_MT_OR_EQU   Lexeme = `MORE_THAN_OR_EQUAL`
	LEXEME_VOID        Lexeme = `VOID`
	LEXEME_TERMINATOR  Lexeme = `TERMINATOR`
)

type Symbol struct {
	Symbol string
	Len    int
	Lexeme Lexeme
}

// IsWordTerminal returns true if the terminal symbol (rune) is allowed within
// a keyword or identifier.
func IsWordTerminal(ru rune) bool {
	return ru == '_' || unicode.IsLetter(ru)
}

func FindKeywordLexeme(kw string) Lexeme {
	switch kw {
	case `F`:
		return LEXEME_FUNC
	case `DO`:
		return LEXEME_DO
	case `END`:
		return LEXEME_END
	case `TRUE`:
		return LEXEME_BOOL
	case `FALSE`:
		return LEXEME_BOOL
	}

	return LEXEME_UNDEFINED
}

const (
	STRING_SYMBOL_START    string = "`"
	STRING_SYMBOL_END      string = "`"
	TEMPLATE_SYMBOL_START  string = `"`
	TEMPLATE_SYMBOL_ESCAPE string = `\`
	TEMPLATE_SYMBOL_END    string = `"`
)

const (
	SYMBOL_COMMENT_START    string = "//"
	SYMBOL_FRACTIONAL_DELIM string = "."
)

func LoneSymbols() []Symbol {
	return []Symbol{
		Symbol{`:=`, 2, LEXEME_ASSIGN},
		Symbol{`->`, 2, LEXEME_RETURNS},
		Symbol{`(`, 1, LEXEME_OPEN_PAREN},
		Symbol{`)`, 1, LEXEME_CLOSE_PAREN},
		Symbol{`[`, 1, LEXEME_OPEN_GUARD},
		Symbol{`]`, 1, LEXEME_CLOSE_GUARD},
		Symbol{`{`, 1, LEXEME_OPEN_LIST},
		Symbol{`}`, 1, LEXEME_CLOSE_LIST},
		Symbol{`,`, 1, LEXEME_DELIM},
		Symbol{`_`, 1, LEXEME_VOID},
		Symbol{`;`, 1, LEXEME_TERMINATOR},
		Symbol{`@`, 1, LEXEME_SPELL},
		Symbol{`~`, 1, LEXEME_NOT},
		Symbol{`¬`, 1, LEXEME_NOT},
		Symbol{`+`, 1, LEXEME_ADD},
		Symbol{`-`, 1, LEXEME_SUBTRACT},
		Symbol{`*`, 1, LEXEME_MULTIPLY},
		Symbol{`/`, 1, LEXEME_DIVIDE},
		Symbol{`%`, 1, LEXEME_REMAINDER},
		Symbol{`&`, 1, LEXEME_AND},
		Symbol{`|`, 1, LEXEME_OR},
		Symbol{`=`, 1, LEXEME_EQU},
		Symbol{`#`, 1, LEXEME_NEQ},
		Symbol{`<=`, 2, LEXEME_LT_OR_EQU},
		Symbol{`=>`, 2, LEXEME_MT_OR_EQU},
		Symbol{`<`, 1, LEXEME_LT},
		Symbol{`>`, 1, LEXEME_MT},
	}
}
