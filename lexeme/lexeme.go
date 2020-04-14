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
// TODO: Some of the const lexeme names are not very meaningful or accurate,
//			 consider making them more precise.
package lexeme

// Lexeme represents a the type of a token. Each lexeme may have multiple
// representations but usually just one.
type Lexeme string

// Enumeration of all possible Lexemes.
const (
	LEXEME_UNDEFINED Lexeme = ``
	LEXEME_ANY       Lexeme = `ANY`
	LEXEME_ANOTHER   Lexeme = `ANOTHER`
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
	LEXEME_PAREN_OPEN  Lexeme = `PAREN_OPEN`
	LEXEME_PAREN_CLOSE Lexeme = `PAREN_CLOSE`
	LEXEME_LIST_OPEN   Lexeme = `LIST_OPEN`
	LEXEME_LIST_CLOSE  Lexeme = `LIST_CLOSE`
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
