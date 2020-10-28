# Versions

## v0.1.0

```
x       := 1
x, y, z := true, 1, "Scarlet"
x       := (1 + 2) * 3
```

### New

- Assignments
	- Single:             `x := 1`
	- Multiple:           `x, y, z := true, 1, "abc"`
- Literals
	- Boolean:            `true`, `false`
	- Numbers:            `1`, `123.456`, `99999999999999999999999999999`
	- Strings:            `"Scarlet"`
- Arithmetic operations
	- Addition:           `1 + 1`
	- Subtraction:        `2 - 1`
	- Multiplication:     `3 * 2`
	- Division:           `4 / 2`
	- Remainder:          `5 % 3`
- Logical operations
	- And:                `true && false`
	- Or:                 `true || false`
- Comparison operations
	- Less than:          `1 < 2`
	- More than:          `2 > 1`
	- Less than or equal: `1 <= 2`
	- More than or equal: `2 >= 1`
	- Equal:              `1 == 1`, `1 == "abc"`
	- Not equal:          `1 != 2`, `1 != "abc"`
- Parameters:           `(1)`, `(1 + 2) * 3`
- Complex expressions:  `(1 + 2 * (3 - 1) == 5) < 2 * 2 * 2 `

## Notes

- Numbers have arbitrary precision

Precedence of operators from highest to lowest:

1. `(`, `)`
2. `*`, `/`, `%`
3. `+`, `-`
4. `<`, `>`, `<=`, `>=`
5. `==`, `!=`
6. `&&`
7. `||`
