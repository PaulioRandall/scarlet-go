# Scribbles

## Current Syntax

```
x       := 1
x, y, z := true, 1, "Scarlet"
x       := (1 + 2) * 3

# Operators & Precedence
first   := "(1)"
second  := "2 * 3 / 4 % 5"
third   := "5 + 6 - 7"
fourth  := "8 < 9 > 10 <= 11 >= 12"
fifth   := "13 == 14 != 15"
sixth   := "true && false"
seventh := "true || false"
```

## Potential Features

```
# Get the program arguments
x := @Args()

# Does a variable exist
x := @Exists("variable_name")

# Find the length of any value whose type has a length
x := @Len(value)

# Stringify a value of any type
x := @Str(value)

# Exit the script with an error message
@Panic(exitCode, message)

# Take a slice of a string
x := @str.Slice(s, startIdx, endIdx)

# Get a specific UTF-8 char from the string, i.e rune
x := @str.Char(s, idx)

# Test if a string has a prefix
x := @str.StartsWith(s, prefix)

# Test if a string has a suffix
x := @str.EndsWith(s, suffix)

# Get the index of a specific UTF-8 char within the string
x := @str.IndexOf(haystack, needle)

# Join two strings together
x := @str.Join("abc", "xyz")

# Parse a string as a number
x, e := @str.ParseNum(number)

# Parse a string as a bool
x, e := @str.ParseBool(bool)

# create a new list
list := @list.New(
  1,
  2,
  3,
)

# Set the value of a list item
@list.Set(list, idx, newValue)

# Get the value of an item in the list
x := @list.Get(list, idx)

# Add an item to the front of a list
@list.Push(list, val)

# Add an item (or another list) to the end of a list
@list.Append(list, item)

# Remove an item from the front of a list
x := @list.Pop(list)

# Remove an item from the end of a list
x := @list.Take(list)

# Take a slice of a list
x := @list.Slice(list, startIdx, endIdx)

# Determine if an index is within a lists range
x := @list.InRange(list, idx)

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