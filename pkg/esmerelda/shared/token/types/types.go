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
	GEN_REDUNDANT:   `REDUNDANT`,
	GEN_TERMINATOR:  `TERMINATOR`,
	GEN_LITERAL:     `LITERAL`,
	GEN_IDENTIFIER:  `IDENTIFIER`,
	GEN_SPELL:       `SPELL`,
	GEN_PARAMS:      `PARAMETERS`,
	GEN_DELIMITER:   `DELIMITER`,
	GEN_PARENTHESIS: `PARENTHESIS`,
}

var subTypes = map[SubType]string{
	SUB_TERMINATOR:  `TERMINATOR`,
	SUB_NEWLINE:     `NEWLINE`,
	SUB_WHITESPACE:  `WHITESPACE`,
	SUB_COMMENT:     `COMMENT`,
	SUB_BOOL:        `BOOL`,
	SUB_NUMBER:      `NUMBER`,
	SUB_STRING:      `STRING`,
	SUB_IDENTIFIER:  `IDENTIFIER`,
	SUB_VOID:        `VOID`,
	SUB_VALUE_DELIM: `VALUE_DELIM`,
	SUB_PAREN_OPEN:  `PAREN_OPEN`,
	SUB_PAREN_CLOSE: `PAREN_CLOSE`,
}

func MaxTypeName() int {

	var max int

	for _, v := range genTypes {
		if len([]rune(v)) > max {
			max = len([]rune(v))
		}
	}

	for _, v := range subTypes {
		if len([]rune(v)) > max {
			max = len([]rune(v))
		}
	}

	return max
}
