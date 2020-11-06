# Versions

## v0.3.0

```
_ := 1 + 2
e := "error"
@Print(@Len(e))
n, e := @ParseNum("123") # 'e' unbound
@Println(n)
@PrintScope()
@Exit(0)
```

- Ignore expression result
- Variable unbinding
- Spells
  - `@Exit(code)` Stops execution of the scroll with the specified exit code
  - `@Print(value...)` Prints the values to terminal in the order provided
  - `@Println(value...)` Same as '@Print' but appends a linefeed
  - `x := @Len(value)` Finds the length of any value whose type has a length
  - `x, e := @ParseNum(number)` Parses a string as a number returning an error message if failed
  - `@PrintScope()` Prints all variables available within the current scope

## v0.2.0

```
x := y
x := y * (3 + z)
```

### API Additions

- Identifier as a term: `x := y`, `x * (y - 1)`

### Notes & Other Changes

- Rewrote large portions of the code base to be simpler and easier to modify.
- Reorganised and amalgamated packages.

## v0.1.0

```
x       := 1
x, y, z := true, 1, "Scarlet"
x       := (1 + 2) * 3
```

### API Additions

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

### Notes & Other Changes

- Numbers have arbitrary precision.

Precedence of operators from highest to lowest:

1. `(`, `)`
2. `*`, `/`, `%`
3. `+`, `-`
4. `<`, `>`, `<=`, `>=`
5. `==`, `!=`
6. `&&`
7. `||`
