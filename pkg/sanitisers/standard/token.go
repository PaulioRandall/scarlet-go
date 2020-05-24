package standard

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type tok struct {
	m Morpheme
	v string
	l int
	c int
}

func (tk tok) Morpheme() Morpheme {
	return tk.m
}

func (tk tok) Value() string {
	return tk.v
}

func (tk tok) Line() int {
	return tk.l
}

func (tk tok) Col() int {
	return tk.c
}
