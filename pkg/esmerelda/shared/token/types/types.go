package types

type GenType int // General type
type SubType int

func (gen GenType) String() string {
	return genTypes[gen]
}

func (sub SubType) String() string {
	return subTypes[sub]
}

const (
	GEN_UNDEFINED GenType = 0
	SUB_UNDEFINED SubType = 0
	GEN_ANY       GenType = iota
	SUB_ANY       SubType = iota
	// ------------------
	GEN_REDUNDANT  GenType = iota
	SUB_WHITESPACE SubType = iota
	SUB_COMMENT
	// ------------------
	GEN_TERMINATOR GenType = iota
	SUB_TERMINATOR SubType = iota
	SUB_NEWLINE
	// ------------------
	GEN_LITERAL GenType = iota
	SUB_BOOL    SubType = iota
	SUB_NUMBER
	SUB_STRING
	// ------------------
	GEN_IDENTIFIER GenType = iota
	SUB_IDENTIFIER SubType = iota
	SUB_VOID
	// ------------------
	GEN_SPELL GenType = iota
	// ------------------
	GEN_PARAMS GenType = iota
	// ------------------
	GEN_DELIMITER   GenType = iota
	SUB_VALUE_DELIM SubType = iota
	// ------------------
	GEN_PARENTHESIS GenType = iota
	SUB_PAREN_OPEN  SubType = iota
	SUB_PAREN_CLOSE
)

var genTypes = map[GenType]string{
	GEN_REDUNDANT:   `redundant`,
	GEN_TERMINATOR:  `terminator`,
	GEN_LITERAL:     `literal`,
	GEN_IDENTIFIER:  `identifier`,
	GEN_SPELL:       `spell`,
	GEN_PARAMS:      `parameters`,
	GEN_DELIMITER:   `delimiter`,
	GEN_PARENTHESIS: `parenthesis`,
}

var subTypes = map[SubType]string{
	SUB_TERMINATOR:  `terminator`,
	SUB_NEWLINE:     `newline`,
	SUB_WHITESPACE:  `whitespace`,
	SUB_COMMENT:     `comment`,
	SUB_BOOL:        `bool`,
	SUB_NUMBER:      `number`,
	SUB_STRING:      `string`,
	SUB_IDENTIFIER:  `identifier`,
	SUB_VOID:        `void`,
	SUB_VALUE_DELIM: `value_delim`,
	SUB_PAREN_OPEN:  `paren_open`,
	SUB_PAREN_CLOSE: `paren_close`,
}
