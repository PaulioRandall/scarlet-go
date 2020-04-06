package token

// Kind represents a token type.
type Kind string

const (
	KIND_UNDEFINED Kind = ``
	// ------------------
	KIND_SOF         Kind = `SOF`
	KIND_EOF         Kind = `EOF`
	KIND_COMMENT     Kind = `COMMENT`
	KIND_WHITESPACE  Kind = `WHITESPACE`
	KIND_NEWLINE     Kind = `NEWLINE`
	KIND_FUNC        Kind = `FUNCTION_KEYWORD`
	KIND_MATCH       Kind = `MATCH_BLOCK_KEYWORD`
	KIND_INLINE      Kind = `INLINE_BLOCK_IMPLIED`
	KIND_ID          Kind = `ID`
	KIND_DELIM       Kind = `DELIMITER`
	KIND_ASSIGN      Kind = `ASSIGNMENT`
	KIND_RETURNS     Kind = `RETURNS`
	KIND_DO          Kind = `BLOCK_OPEN`
	KIND_END         Kind = `BLOCK_CLOSE`
	KIND_OPEN_PAREN  Kind = `PAREN_OPEN`
	KIND_CLOSE_PAREN Kind = `PAREN_CLOSE`
	KIND_OPEN_GUARD  Kind = `GUARD_OPEN`
	KIND_CLOSE_GUARD Kind = `GUARD_CLOSE`
	KIND_OPEN_LIST   Kind = `LIST_OPEN`
	KIND_CLOSE_LIST  Kind = `LIST_CLOSE`
	// TODO: kind these
	SPELL      Kind = `SPELL`
	STR        Kind = `STR`
	TEMPLATE   Kind = `TEMPLATE`
	INT        Kind = `INT`
	REAL       Kind = `REAL`
	BOOL       Kind = `BOOL`
	NOT        Kind = `NOT`
	ADD        Kind = `ADD`
	SUBTRACT   Kind = `SUBTRACT`
	MULTIPLY   Kind = `MULTIPLY`
	DIVIDE     Kind = `DIVIDE`
	MOD        Kind = `MOD`
	AND        Kind = `AND`
	OR         Kind = `OR`
	EQU        Kind = `EQUAL`
	NEQ        Kind = `NOT_EQUAL`
	LT         Kind = `LESS_THAN`
	LT_OR_EQU  Kind = `LESS_THAN_OR_EQUAL`
	MT         Kind = `MORE_THAN`
	MT_OR_EQU  Kind = `MORE_THAN_OR_EQUAL`
	VOID       Kind = `VOID`
	TERMINATOR Kind = `TERMINATOR`
)

const (
	KEYWORD_FUNCTION    string = `F`
	KEYWORD_BLOCK_START string = `DO`
	KEYWORD_BLOCK_END   string = `END`
	KEYWORD_TRUE        string = `TRUE`
	KEYWORD_FALSE       string = `FALSE`
)

// KeywordToKind maps a non-terminal keyword to a token kind.
func KeywordToKind(nonTerminal string) Kind {

	switch nonTerminal {
	case KEYWORD_FUNCTION:
		return KIND_FUNC
	case KEYWORD_BLOCK_START:
		return KIND_DO
	case KEYWORD_BLOCK_END:
		return KIND_END
	case KEYWORD_TRUE, KEYWORD_FALSE:
		return BOOL
	}

	return KIND_UNDEFINED
}

const (
	NEWLINE_LF   string = "\n"
	NEWLINE_CRLF string = "\r\n"
)

const (
	STRING_SYMBOL_START    string = "`"
	STRING_SYMBOL_END      string = "`"
	TEMPLATE_SYMBOL_START  string = `"`
	TEMPLATE_SYMBOL_ESCAPE string = `\`
	TEMPLATE_SYMBOL_END    string = `"`
)

const (
	SYMBOL_COMMENT_START    string = "/"
	SYMBOL_FRACTIONAL_DELIM string = "."
)

type LoneSymbol struct {
	Symbol string
	Len    int
	Kind   Kind
}

func LoneSymbols() []LoneSymbol {
	return []LoneSymbol{
		LoneSymbol{`:=`, 2, KIND_ASSIGN},
		LoneSymbol{`->`, 2, KIND_RETURNS},
		LoneSymbol{`(`, 1, KIND_OPEN_PAREN},
		LoneSymbol{`)`, 1, KIND_CLOSE_PAREN},
		LoneSymbol{`[`, 1, KIND_OPEN_GUARD},
		LoneSymbol{`]`, 1, KIND_CLOSE_GUARD},
		LoneSymbol{`{`, 1, KIND_OPEN_LIST},
		LoneSymbol{`}`, 1, KIND_CLOSE_LIST},
		LoneSymbol{`,`, 1, KIND_DELIM},
		LoneSymbol{`_`, 1, VOID},
		LoneSymbol{`;`, 1, TERMINATOR},
		LoneSymbol{`@`, 1, SPELL},
		LoneSymbol{`~`, 1, NOT},
		LoneSymbol{`¬`, 1, NOT},
		LoneSymbol{`+`, 1, ADD},
		LoneSymbol{`-`, 1, SUBTRACT},
		LoneSymbol{`*`, 1, MULTIPLY},
		LoneSymbol{`/`, 1, DIVIDE},
		LoneSymbol{`%`, 1, MOD},
		LoneSymbol{`&`, 1, AND},
		LoneSymbol{`|`, 1, OR},
		LoneSymbol{`=`, 1, EQU},
		LoneSymbol{`#`, 1, NEQ},
		LoneSymbol{`<=`, 2, LT_OR_EQU},
		LoneSymbol{`=>`, 2, MT_OR_EQU},
		LoneSymbol{`<`, 1, LT},
		LoneSymbol{`>`, 1, MT},
	}
}
