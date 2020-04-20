// token package was created to define the range of possible token types and
// house the Token structure, which is core to all components.
//
// Key decisions: Can you think of any decisions that need explaining?
//
// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.
package token

// TokenType represents the type of a token. Each may have multiple lexemes
// but usually just the one (lemma).
type TokenType string

const (
	UNDEFINED TokenType = ``
	// ------------------
	ANY     TokenType = `ANY`
	ANOTHER TokenType = `ANOTHER`
	// ------------------
	EOF TokenType = `EOF`
	// ------------------
	COMMENT            TokenType = `COMMENT`
	WHITESPACE         TokenType = `WHITESPACE`
	NEWLINE            TokenType = `NEWLINE`
	FUNC               TokenType = `FUNCTION`
	ID                 TokenType = `ID`
	DELIM              TokenType = `DELIMITER`
	ASSIGN             TokenType = `ASSIGNS`
	RETURNS            TokenType = `RETURNS`
	MATCH_OPEN         TokenType = `MATCH`
	BLOCK_OPEN         TokenType = `DO`
	BLOCK_CLOSE        TokenType = `END`
	PAREN_OPEN         TokenType = `PAREN_OPEN`
	PAREN_CLOSE        TokenType = `PAREN_CLOSE`
	LIST               TokenType = `LIST`
	LIST_OPEN          TokenType = `LIST_OPEN`
	LIST_CLOSE         TokenType = `LIST_CLOSE`
	GUARD_OPEN         TokenType = `GUARD_OPEN`
	GUARD_CLOSE        TokenType = `GUARD_CLOSE`
	SPELL              TokenType = `SPELL`
	STRING             TokenType = `STRING`
	TEMPLATE           TokenType = `TEMPLATE`
	NUMBER             TokenType = `NUMBER`
	BOOL               TokenType = `BOOL`
	ADD                TokenType = `ADD`
	SUBTRACT           TokenType = `SUBTRACT`
	MULTIPLY           TokenType = `MULTIPLY`
	DIVIDE             TokenType = `DIVIDE`
	REMAINDER          TokenType = `REMAINDER`
	AND                TokenType = `AND`
	OR                 TokenType = `OR`
	EQUAL              TokenType = `EQUAL`
	NOT_EQUAL          TokenType = `NOT_EQUAL`
	LESS_THAN          TokenType = `LESS_THAN`
	LESS_THAN_OR_EQUAL TokenType = `LESS_THAN_OR_EQUAL`
	MORE_THAN          TokenType = `MORE_THAN`
	MORE_THAN_OR_EQUAL TokenType = `MORE_THAN_OR_EQUAL`
	VOID               TokenType = `VOID`
	TERMINATOR         TokenType = `TERMINATOR`
)
