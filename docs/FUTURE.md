# Scribbles

## [v0.7.0]

```
# Create a new initialised map 
map <- @NewMap(
  "one", 1,
  "two", 2,
  "three", 3,
)

# Map a value to a key
@Set(map, key, value)

# Get the value of map entry using its key
v <- @Get(map, key)

# Remove a map entry
v <- @Del(map, key)

# Get a list of all keys in a map
list <- @Keys(map)

# Get a list of all values in a map
list <- @Values(map)
```

## [v0.8.0+] Potential Features

- Manual/documentation
- Native lists & maps
- Useful error messages

```

# Write to standard output, a space is placed between each printed item 
<< "abc", "efg"

# Write to standard output, a newline is placed after each printed item 
<<< "abc", "efg"

# Get the program arguments
x := @Args()

# Iterate a list
x := @list.Foreach(list, F(i, value, more) {
  ...
})

# Iterate a map
x := @map.Foreach(map, F(key, value) {
  ...
})

# Example expression function
add := E(a, b, c) a + b + c
x := add(1, 2, 3)

# Example function
f := F(a, b -> x, y) {
  ...
}
x, y := f(1, 2)

# Exit the current function
exit F

# Non-function Blocks (as spell parameter only): A block of code that is passed
# to a spell but is run within the context of the calling scope
@if(condition, {
  ...
})

# Definitions (AKA globals constants): Definable only at the root scope of a
# script, i.e. not within functions
def x := 1
def x, y, z := 1, 2, 3
def f := F(a, b) {
  ...
}

# While loop
loop more := true [more] {
  ...
}

# Ranged loop
loop i := init [i < size] {
  ...
}

# Exit the current loop
exit loop

# Native Foreach Loops
loop list -> i, value, more {
  ...
}

# Watch Blocks: Watches a variable and exits the block if a change occurs at
# the end of any statement
watch e {
  ...
  e := "error"
  ...
}

# Template Strings & Spells
a, op, b, eql, c := 1, "+", 2, "=", 3
x := @Fmt("{a} {op} {b} {eql} {c}")     # "1 + 2 = 3"

a, b, c := 1.1, 2.22, 3.333
x := @Fmt("{a, .2} + {b, .2} = {c, .2}") # "1.10 + 2.22 + 3.33"

# Catch Spell: Executes the code in the block and returns the error if there is
# a panic
e := @Catch({
  ...
})
```
