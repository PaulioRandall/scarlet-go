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
	SU_UNDEFINED  SubType = 0
	GEN_ANY       GenType = iota
	SU_ANY        SubType = iota
	// ------------------
	GEN_WHITESPACE GenType = iota
	// ------------------
	GEN_TERMINATOR GenType = iota
	SU_TERMINATOR  SubType = iota
	SU_NEWLINE
	// ------------------
	GEN_LITERAL GenType = iota
	SU_BOOL     SubType = iota
	SU_NUMBER
	SU_STRING
	// ------------------
	GEN_IDENTIFIER GenType = iota
	SU_IDENTIFIER  SubType = iota
	SU_VOID
	// ------------------
	GEN_SPELL GenType = iota
	// ------------------
	GEN_PARAMS GenType = iota
	// ------------------
	GEN_DELIMITER  GenType = iota
	SU_VALUE_DELIM SubType = iota
	// ------------------
	GEN_PARENTHESIS GenType = iota
	SU_PAREN_OPEN   SubType = iota
	SU_PAREN_CLOSE
)

var genTypes = map[GenType]string{
	GEN_WHITESPACE:  `whitespace`,
	GEN_TERMINATOR:  `terminator`,
	GEN_LITERAL:     `literal`,
	GEN_IDENTIFIER:  `identifier`,
	GEN_SPELL:       `spell`,
	GEN_PARAMS:      `parameters`,
	GEN_DELIMITER:   `delimiter`,
	GEN_PARENTHESIS: `parenthesis`,
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
