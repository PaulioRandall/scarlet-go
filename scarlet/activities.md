
# MAYBE: Init output variables
### Examples
```
F(list, ^sum: 0) {
	LOOP i, v, m <- list {
		sum: sum + v
	}
}
```


# ERROR: Terminator check required
- This passes during parsing but shouldn't 
- `x: 1 2`


# Allow voids as assignment targets
Void assignment targets ignore the result of an expression, useful for indicating that a result is not needed


# How can dependencies be reduced?
Better definition and use of interfaces will make for more segregated code, which will be easier to maintain.


# Exit script early
### Examples
- `EXIT SCRIPT`


# Exit loop early
### Examples
- `EXIT LOOP`


# Exit guard early
### Examples
- `EXIT GUARD`


# Exit function early
### Examples
- `EXIT F`


# Enhance loops
### Allow
1. LOOP i := 0 [i < 5]
2. LOOP i, x := 0, 1 [i < 5]


# Add inbuilt functions
### Print function
```
Prints the args (variable length) to console
@P(...)
```

### Print line function
```
Prints the args (variable length) to console, and appends a linefeed
@PL(...)
```

### Examples
- `@P("x: ", 1 + 2, "; ")`
- `@PL(list)`


# Can Keywords be case-insensitive?
Maybe they should be lower case, except for F?


# Template strings
### Example
- `s := "alpha = {list[0]}, beta = {list[1]}"`


# Expression functions
Expression functions have a single expression as their body. The result of the expression is returned.

### Examples
- `increment := E(n) n + 1` 
- `expr := E(a, b) [a > b] @P("A > B")`


# Key-value pairs
### Examples
- `p := "key": "value"`
- `k, v := p[K], p[V]`


# Simple directory navigation spells
### Spells
- `@cd("./scarlet-go")`
- `@pushd("./scarlet-go")`
- `@popd()`


# Quick shell commands
### Examples 
- `$ "go build -o" filename "scarlet.go"`


# Better token type naming


# Everything that can fail with ASSIGN parsing


# Everything that can fail with FUNC parsing


# Everything that can fail with GUARD parsing


# Everything that can fail with MAtCH parsing


# Everything that can fail with LOOP parsing


# Everything that can fail with EXPRESSION parsing


# Everything that can fail with LIST parsing
