package statement

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type builder struct {
	sb strings.Builder
}

func (b *builder) add(indent int, s string) {

	for _, ru := range s {
		b.sb.WriteRune(ru)

		if ru == '\n' {
			for i := 0; i < indent; i++ {
				b.sb.WriteRune('\t')
			}
		}
	}
}

func (b *builder) addToken(indent int, tk Token) {
	b.add(indent, ToString(tk))
}

func (b *builder) newline(indent int) {
	b.add(indent, "\n")
}

func (b *builder) String() string {
	return b.String()
}

func (b *builder) print() {
	fmt.Println(b.String())
	fmt.Println()
}
