package statement

import (
	"fmt"
	"strings"

	. "github.com/PaulioRandall/scarlet-go/pkg/token"
)

type builder strings.Builder

func (b *builder) add(s string) {
	sb := strings.Builder(*b)
	sb.WriteString(s)
}

func (b *builder) addToken(tk Token) {
	b.add(ToString(tk))
}

func (b *builder) indent(indent int) {
	b.add(strings.Repeat("\t", indent))
}

func (b *builder) newline() {
	b.add("\n")
}

func (b *builder) String() string {
	return b.String()
}

func (b *builder) print() {
	fmt.Println(b.String())
	fmt.Println()
}
