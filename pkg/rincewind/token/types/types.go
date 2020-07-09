package types

type GenType int // General type
type SubType int

func (ge GenType) String() string {
	return genTypes[ge]
}

func (su SubType) String() string {
	return subTypes[su]
}

const (
	GE_UNDEFINED GenType = 0
	SU_UNDEFINED SubType = 0
	GE_ANY       GenType = iota
	SU_ANY       SubType = iota
	// ------------------
	GE_WHITESPACE GenType = iota
	// ------------------
	GE_TERMINATOR GenType = iota
	SU_TERMINATOR SubType = iota
	SU_NEWLINE
	// ------------------
	GE_LITERAL GenType = iota
	SU_BOOL    SubType = iota
	SU_NUMBER
	SU_STRING
	// ------------------
	GE_IDENTIFIER GenType = iota
	SU_IDENTIFIER SubType = iota
	SU_VOID
	// ------------------
	GE_SPELL GenType = iota
	// ------------------
	GE_PARAMS GenType = iota
	// ------------------
	GE_DELIMITER   GenType = iota
	SU_VALUE_DELIM SubType = iota
	// ------------------
	GE_PARENTHESIS GenType = iota
	SU_PAREN_OPEN  SubType = iota
	SU_PAREN_CLOSE
)

var genTypes = map[GenType]string{
	GE_WHITESPACE:  `whitespace`,
	GE_TERMINATOR:  `terminator`,
	GE_LITERAL:     `literal`,
	GE_IDENTIFIER:  `identifier`,
	GE_SPELL:       `spell`,
	GE_PARAMS:      `parameters`,
	GE_DELIMITER:   `delimiter`,
	GE_PARENTHESIS: `parenthesis`,
}

var subTypes = map[SubType]string{
	SU_TERMINATOR:  `terminator`,
	SU_NEWLINE:     `newline`,
	SU_BOOL:        `bool`,
	SU_NUMBER:      `number`,
	SU_STRING:      `string`,
	SU_IDENTIFIER:  `identifier`,
	SU_VOID:        `void`,
	SU_VALUE_DELIM: `value_delim`,
	SU_PAREN_OPEN:  `paren_open`,
	SU_PAREN_CLOSE: `paren_close`,
}
