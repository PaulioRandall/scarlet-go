
# MAYBE: Parameterless functions as function arguments could be written
- `{ @pl("Scarlet") }`


# MAYBE: Parameterless expression functions could be written
- `E{a+1}`


# MAYBE: Loops could be rewritten as spells?
- Loops could then be removed from the language
### Examples

```
@For({i: 1}, E{i<=5}, {
	@pl(i)
})
```

```
@While(E{i<=5}, {
	@pl(i)
})
```


# Write spells for list appending & prepending
### Examples
- `e: @append(list, x)`
- `e: @prepend(x, list)`


# Write spell for slicing lists
- `new_list, e: @slice(list, 2, 4)`


# Write spell for getting length of list or string
- `x, e: @len(list)`


# Write spell for dividing safely
- `x, e: @div(a, b)`


# Write spell to parse number
- `n, e: @parseNum(numStr)`


# Write spell to process file data
- `e: @file(filePath, f)`


# Write spell to read all data from a file 
- `s, e: @readFile(filePath)`


# Allow functions that return a function to be called
### Examples
- `x: f()()`


# ERROR: Terminator checks required
- This passes during parsing but shouldn't 
- `x: 1 2`


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


# Can Keywords be case-insensitive?
Maybe they should be lower case, except for `F` and `E`?


# Template strings
### Example
- `s := "alpha = {list[0]}, beta = {list[1]}"`


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
