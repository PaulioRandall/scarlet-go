// lexeme package was created to define the range of possible types of tokens,
// or lexemes as they have been named here. The package also defines the actual
// terminal and non-terminal symbols used in the language.
//
// Key decisions:
// 1. To be honest, I really don't know if it was a good idea to use the word
// lexeme in the manner I have. Usage in other literature doesn't seem
// consistent with each other so I'm just going to run with what I've got until
// I have a better understanding and alternative word.
// 2. I centralised the definition of terminal and non-terminal symbols here
// so they could be modified and tweaked easily. I'm not sure that was a good
// idea either. It is highly coupled with the logic which parses tokens so I'm
// thinking about moving them back to the scanner package, but perhaps defined
// in their own file to maintain ease of modification.
//
// TODO: Consider moving the terminal and non-terminal definitions back to the
// 			 scanner package but in their own file for ease of readability and
// 			 modification.
// TODO: Some of the const lexeme names are not very meaningful or accurate,
//			 consider making them more precise.
// TODO: Scanning of keywords and symbols could be combined with identifier
//       scanning being attempted after symbol and keyword scanning. The
//       advantage is that keywords and lone symbols would be interchangable,
//			 which I argue they should be. It would allow `DO` `END` blocks to be
//			 redefined using `{` `}` if need be, with changes to list definitions
//			 as well of course.
package lexeme

import (
	"unicode"
)

// Lexeme represents a the type of a token. Each lexeme may have multiple
// representations but usually just one.
type Lexeme string

// Enumeration of all possible Lexemes.
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

// IsWordTerminal returns true if the terminal symbol (rune) is allowed within
// a keyword or identifier.
func IsWordTerminal(ru rune) bool {
	return ru == '_' || unicode.IsLetter(ru)
}

// Defining the non-terminals used for delimiting strings.
const (
	STRING_SYMBOL_START    string = "`"
	STRING_SYMBOL_END      string = "`"
	TEMPLATE_SYMBOL_START  string = `"`
	TEMPLATE_SYMBOL_ESCAPE string = `\`
	TEMPLATE_SYMBOL_END    string = `"`
)

// Defining non-terminals that have no other sensible home.
const (
	SYMBOL_COMMENT_START    string = "//"
	SYMBOL_FRACTIONAL_DELIM string = "."
)

// Symbol represents a non-terminal symbol that is not a keyword, yet has a
// lexeme of it's own.
type Symbol struct {
	Symbol   string
	Len      int
	Lexeme   Lexeme
	ScanFunc func([]rune) (Token, bool)
}

// Symbols returns an array of all possible non-terminal symbols that are not
// keywords, yet have a lexeme of their own. Longest and highest priority
// symbols should be at the beginning of the array to ensure the correct token
// is scanned.
func Symbols() []Symbol {
	return []Symbol{
		Symbol{`FALSE`, 5, LEXEME_BOOL, nil},
		Symbol{`TRUE`, 4, LEXEME_BOOL, nil},
		Symbol{`END`, 3, LEXEME_END, nil},
		Symbol{`DO`, 2, LEXEME_DO, nil},
		Symbol{`F`, 1, LEXEME_FUNC, nil},
		Symbol{`:=`, 2, LEXEME_ASSIGN, nil},
		Symbol{`->`, 2, LEXEME_RETURNS, nil},
		Symbol{`<=`, 2, LEXEME_LT_OR_EQU, nil},
		Symbol{`=>`, 2, LEXEME_MT_OR_EQU, nil},
		Symbol{`(`, 1, LEXEME_OPEN_PAREN, nil},
		Symbol{`)`, 1, LEXEME_CLOSE_PAREN, nil},
		Symbol{`[`, 1, LEXEME_OPEN_GUARD, nil},
		Symbol{`]`, 1, LEXEME_CLOSE_GUARD, nil},
		Symbol{`{`, 1, LEXEME_OPEN_LIST, nil},
		Symbol{`}`, 1, LEXEME_CLOSE_LIST, nil},
		Symbol{`,`, 1, LEXEME_DELIM, nil},
		Symbol{`_`, 1, LEXEME_VOID, nil},
		Symbol{`;`, 1, LEXEME_TERMINATOR, nil},
		Symbol{`@`, 1, LEXEME_SPELL, nil},
		Symbol{`+`, 1, LEXEME_ADD, nil},
		Symbol{`-`, 1, LEXEME_SUBTRACT, nil},
		Symbol{`*`, 1, LEXEME_MULTIPLY, nil},
		Symbol{`/`, 1, LEXEME_DIVIDE, nil},
		Symbol{`%`, 1, LEXEME_REMAINDER, nil},
		Symbol{`&`, 1, LEXEME_AND, nil},
		Symbol{`|`, 1, LEXEME_OR, nil},
		Symbol{`=`, 1, LEXEME_EQU, nil},
		Symbol{`#`, 1, LEXEME_NEQ, nil},
		Symbol{`<`, 1, LEXEME_LT, nil},
		Symbol{`>`, 1, LEXEME_MT, nil},
	}
}
