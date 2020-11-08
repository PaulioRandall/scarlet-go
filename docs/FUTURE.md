# Scribbles

## [v0.5.0]

```
s <- "Happy"
i <- @Index(s, "app")
s <- @Join(s, " days!")
ten <- @Str(10)
@Panic(1, "Meh")
```

- Assignment symbol changed to `<-`
- `i <- @Index(con, s)` Returns the index of an item within a container or -1 if the item doesn't exists
- `con <- @Join(con, con)` Joins two containers 
- `s <- @Str(value)` Stringify a value
- `@Panic(exitCode, message)` Exit the scroll after printing an error message

## [v0.6.0+] Potential Features

- Manual/documentation
- Native lists & maps

```
# Write to standard output, a space is placed between each printed item 
<< "abc", "efg"

# Write to standard output, a newline is placed after each printed item 
<<< "abc", "efg"

# Get the program arguments
x := @Args()

# Does a variable exist
x := @Exists("variable_name")

# Iterate a list
x := @list.Foreach(list, F(i, value, more) {
  ...
})

# Create a new map
map := @map.New(
  1, "one",
  2, "two",
  3, "three",
)

# Map a value to a key
@map.Set(map, key, value)

# Get the value of map entry using its key
x := @map.Get(map, key)

# Remove a map entry
x := @map.Del(key)

# Get a list of all keys in a map
x := @map.Keys(map)

# Get a list of all values in a map
x := @map.Values(map)

# Test if a key exists within a map
x := @map.Exists(key)

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

# Test if a variable exists
x := y?

# When Block: A form of match block or switch
when {
  [x < 0] { // Guard case
    ... 
  }
  [true] { // Default case
    ... 
  }
}

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
x := @fmt.Fmt("{a} {op} {b} {eql} {c}") # "1 + 2 = 3"

a, b, c := 1.1, 2.22, 3.333
x := @Fmt("{a, .2} + {b, .2} = {c, .2}") # "1.10 + 2.22 + 3.33"

# Catch Spell: Executes the code in the block and returns the error if there is
# a panic
e := @Catch({
  ...
})
```
