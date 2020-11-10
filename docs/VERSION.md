# Versions

## [v0.6.0]

```
[2 > 1] @Println("Scarlet!)
[1 < 2] {
  @Println("Scarlet!")
}
when {
  [x < 0] @Println("x is negative")
  [x > 0] @Println("x is positive")
  [true] {
    @Println("x is zero")
  }
}
```

- `[cond] stmt` Guarded or conditional statement
- `[cond] { stmts }` Guarded or conditional block
- `when { [cond] stmt;... }` Match block where only the body of the first true guard is executed

## [v0.5.0]

```
s <- "Happy"
i <- @Index(s, "app")
s <- @Join(s, " days!")
ten <- @Str(10)
@Panic(1, "Meh")
exists <- err?
```

- Assignment symbol changed to `<-`
- `i <- @Index(con, s)` Returns the index of an item within a container or unbinds the target identifier if not found
- `con <- @Join(con, con)` Joins two containers 
- `s <- @Str(value)` Stringify a value
- `@Panic(exitCode, message)` Exit the scroll after printing an error message
- `ok <- value?` Results in true if a value exists, can be used on identifiers

## [v0.4.0]

```
list : @NewList(1, 2,	true,	"abc")
list : @Slice(list, 1, 3)   # Returns [2, true, "abc"]
str : @At(list, 2)
ok : @InRange(list, 9)      # Returns false
list : @Push(list, 123.456) # Returns [123.456, 2, true, "abc"]
list : @Add(list, "happy")  # Returns [123.456, 2, true, "abc", "happy"]
list : @Set(list, 2, false) # Returns [123.456, 2, false, "abc", "happy"]
list, v : @Del(list, 3)     # Returns [123.456, 2, false, "happy"], "abc"
list, v : @Pop(list)        # Returns [2, false, "happy"], 123.456
list, v : @Take(list)       # Returns [123.456, 2, false], "happy"
```

- All numbers are now float64, arbitrary precision removed
- New list type
  - `list : @NewList(value...)` Creates a new ordered list
- Spells
  - `con : @Slice(con, start, end)` Take a slice of a container
  - `v : @At(con, index)` Gets the value at the specified index
  - `ok : @InRange(con, idx)` Returns true if the index is valid within a container
  - `con : @Push(con, value...)` Pushes values onto the front of a container
  - `con : @Add(con, value...)` Pushes values onto the end of a container
  - `con : @Set(con, index, value)` Sets the value of an item within a container, not usable with strings
  - `con, v : @Del(con, index)` Removes an item from a container
  - `con, v : @Pop(con)` Removes the item at the front of a container
  - `con, v : @Take(con)` Removes the item at the back of a container

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
