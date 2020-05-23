package z_matching

import (
	. "github.com/PaulioRandall/scarlet-go/pkg/z_token"
)

type tok struct {
	k Kind
	m Morpheme
	v string
	l int
	c int
}

func (tk tok) Kind() Kind {
	return tk.k
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
