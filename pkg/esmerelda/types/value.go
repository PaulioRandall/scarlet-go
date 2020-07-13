package types

type Value interface {
	Equal(other Value) bool
	Comparable(other Value) bool
	String() string
}
