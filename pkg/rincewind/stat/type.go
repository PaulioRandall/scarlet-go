package stat

type StatType int

func (st StatType) String() string {
	return types[st]
}

const (
	ST_UNDEFINED StatType = iota
	ST_SPELL_CALL
)

var types map[StatType]string = map[StatType]string{
	ST_UNDEFINED:  ``,
	ST_SPELL_CALL: `Spell-Call`,
}
