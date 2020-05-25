
# ERROR: Terminator check required
- This passes during parsing but shouldn't 
- `x: 1 2`


# Add list-based looping
### Formats
- `LOOP ID DELIM ID DELIM ID updates expression<list> full-block`

### Examples
- `LOOP index, value, hasMore <- list {}`
- `LOOP i, v, m <- f() {}`

### Steps
1. Add `<-` (UPDATES) symbol to scanner
2. Add Iterator struct to statement pkg
3. Add pattern to parser
4. Add execution of Iterator to runtime

### Tests
1. Scanner token: UPDATES `<-`
2. Scanner statement: `LOOP i, v, m <- list {}`
3. Scanner statement: `LOOP i, v, m <- f() {}`
4. Parser statement: `LOOP ID DELIM ID DELIM ID UPDATES ID BLOCK_OPEN ID ASSIGN NUMBER BLOCK_CLOSE`
5. Parser statement: `LOOP ID DELIM ID DELIM ID UPDATES ID PAREN_OPEN PAREN_CLOSE BLOCK_OPEN ID ASSIGN NUMBER BLOCK_CLOSE`


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
