package token

type TokenType string

// TODO: Some of the const token types don't have meaningful or accurate names,
//			 consider improving matters.

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
	FIX                TokenType = `FIX`
	LET                TokenType = `LET`
	ID                 TokenType = `ID`
	DELIM              TokenType = `DELIMITER`
	ASSIGN             TokenType = `ASSIGNS`
	OUTPUT             TokenType = `OUTPUT_PARAM`
	BLOCK_OPEN         TokenType = `NEW_BLOCK`
	BLOCK_CLOSE        TokenType = `END_BLOCK`
	PAREN_OPEN         TokenType = `PAREN_OPEN`
	PAREN_CLOSE        TokenType = `PAREN_CLOSE`
	LIST               TokenType = `LIST`
	MATCH              TokenType = `MATCH`
	GUARD_OPEN         TokenType = `GUARD_OPEN`
	GUARD_CLOSE        TokenType = `GUARD_CLOSE`
	LOOP               TokenType = `LOOP`
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
