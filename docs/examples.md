# Examples & Formats

## Identifiers

#### Formats
- identifier    := "\_"
- identifier    := lower_letter {letter}
- \~letter       := upper_letter | lower_letter | "\_"
- \~upper_letter := "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H" | "I" | "J" | "K" | "L" | "M" | "N" | "O" | "P" | "Q" | "R" | "S" | "T" | "U" | "V" | "W" | "X" | "Y" | "Z"
- \~lower_letter := "a" | "b" | "c" | "d" | "e" | "f" | "g" | "h" | "i" | "j" | "k" | "l" | "m" | "n" | "o" | "p" | "q" | "r" | "s" | "t" | "u" | "v" | "w" | "x" | "y" | "z"

#### Examples
- `_`
- `x`
- `abc`
- `abc_EFG`

## Literals

#### Formats
- literal  := bool | number | string
- bool     := "TRUE" | "FALSE"
- number   := integer ["." integer]
- string   := '"' * Any UTF-8 code-point * '"'
- \~integer := digit {digit}
- \~digit   := "0" | 1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"

#### Examples
- `TRUE`
- `1`
- `1.1`
- `"abc"`

## Lists

#### Formats
- list        := "LIST" "{" [parameters] "}"
- \~parameters := expression {"," expression} [","]

#### Examples
- `LIST {}`
- `LIST {a,b,c}`
- `LIST {a,b,c,}`

```
LIST {
	a,
	b,
	c,
}
```

## List Access

#### Formats
- list_access := identifier "[" num_expression "]"

#### Examples
- `data[2]`
- `data[i]`

## Expressions

#### Formats
- expression      := identifier | literal | list | list_access
- expression      := "-" expression
- bool_expression := bool
- num_expression  := number

#### Examples
- `x`
- `1`
- `LIST {a,b,c}`
- `data[2]`
- `-1`
- `-a`
