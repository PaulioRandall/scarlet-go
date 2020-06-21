package statement

import (
	"fmt"
	"strings"
)

type builder struct {
	strings.Builder
}

func (b *builder) add(indent int, s string) {

	for _, ru := range s {
		b.WriteRune(ru)

		if ru == '\n' {
			for i := 0; i < indent; i++ {
				b.WriteRune('\t')
			}
		}
	}
}

func (b *builder) newline() {
	b.add(0, "\n")
}

func (b *builder) String() string {
	return b.Builder.String()
}

func (b *builder) print() {
	fmt.Println(b.String())
	fmt.Println()
}
