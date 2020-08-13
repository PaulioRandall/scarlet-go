# scarlet-go

Scarlet is my second attempt at creating an interpreter. Scarlett is the name of
the default language.

## 1.0 (Essential Features)

#### Arithmetic Operations

```
1 + 2
1 - 2
1 * 2
1 / 2
1 % 2
1 * (2 + 3)
```

```
(1 + 2) * 3 - 4 / 5 % 6
```

#### Logical Operations

```
true && false
true || false
```

#### Comparison Operations

```
1 < 2
1 > 2
1 <= 2
1 >= 2
1 == 2
1 != 2
```

#### Guarded Statements

```
[x > 1] {
  ...
}
```

```
[x > 1] ...
```

#### While Loops

```
loop [true] {
  ...
}
```

## 1.1

#### Void Assignee

```
_, y := 1, 2
```

#### Native Variable Deletions

```
x := _
```

#### Spell Returns

```
sum := @Sum(1, 2, 3)
```

#### Program Spells

Get the program arguments:
```
x := @Args()
```

Does a variable exist:
```
x := @Exists("variable_name")
```

Find the length of any value whose type has a length:
```
x := @Len(value)
```

Stringify a value of any type:
```
x := @Str(value)
```

Exit the script with an error message:
```
@Panic(exitCode, message)
```

## 1.2 (String Spells)

Take a slice of a string:
```
x := @str.Slice(s, startIdx, endIdx)
```

Get a specific UTF-8 char from the string, i.e rune:
```
x := @str.Char(s, idx)
```

Test if a string has a prefix:
```
x := @str.StartsWith(s, prefix)
```

Test if a string has a suffix:
```
x := @str.EndsWith(s, suffix)
```

Get the index of a specific UTF-8 char within the string:
```
x := @str.IndexOf(haystack, needle)
```

Join two strings together:
```
x := @str.Join("abc", "xyz")
```

Parse a string as a number:
```
x, e := @str.ParseNum(number)
```

Parse a string as a bool:
```
x, e := @str.ParseBool(bool)
```

## 2.0 

#### Functions

```
f := F(a, b -> x, y) {
  ...
}

x, y := f(1, 2)
```

## 3.0 (Lists)

Create a new list:
```
list := @list.New(
  1,
  2,
  3,
)
```

Set the value of a list item:
```
@list.Set(list, idx, newValue)
```

Get the value of an item in the list:
```
x := @list.Get(list, idx)
```

Add an item to the front of a list:
```
@list.Push(list, val)
```

Add an item (or another list) to the end of a list:
```
@list.Append(list, item)
```

Remove an item from the front of a list:
```
x := @list.Pop(list)
```

Remove an item from the end of a list:
```
x := @list.Take(list)
```

Take a slice of a list:
```
x := @list.Slice(list, startIdx, endIdx)
```

Determine if an index is within a lists range:
```
x := @list.InRange(list, idx)
```

Iterate a list:
```
x := @list.Foreach(list, F(i, value, more) {
  ...
})
```

## 3.1 (Maps)

Create a new map:
```
map := @map.New(
  1, "one",
  2, "two",
  3, "three",
)
```

Map a value to a key:
```
@map.Set(map, key, value)
```

Get the value of map entry using its key:
```
x := @map.Get(map, key)
```

Remove a map entry:
```
x := @map.Del(key)
```

Get a list of all keys in a map:
```
x := @map.Keys(map)
```

Get a list of all values in a map:
```
x := @map.Values(map)
```

Test if a key exists within a map:
```
x := @map.Exists(key)
```

Iterate a map:
```
x := @map.Foreach(map, F(key, value) {
  ...
})
```

## 4.0+

These are debatable features that are not really required but might make
programming moderately smoother.

#### Variable Existence

Test if a variable exists.
```
x := y?
```

#### When Blocks

A form of match block or switch.
```
when x {
  1: ...
  2: {
    ...
  }
  [x < 0] { // Guard case
    ... 
  }
  [true] { // Default case
    ... 
  }
}
```

#### Exit Function

Exit the current function:
```
exit F
```

#### Non-function Blocks (as spell parameter only)

A block of code that is passed to a spell but is run within the context of the
calling scope.
```
@if(condition, {
  ...
})
```

#### Definitions

```
def x := 1
```

```
def f := F(a, b) {
  ...
}
```

```
def x, y, z := 1, 2, 3
```

#### Native While & For Loops

Infinite loop:
```
loop [true] {
  ...
}
```

While loop:
```
loop more := true [more] {
  ...
}
```

Ranged loop:
```
loop i := init [i < size] {
  ...
}
```

Exit the current loop:
```
exit loop
```

#### Native Foreach Loops

```
loop list -> i, value, more {
  ...
}
```

#### Watch Blocks

Watches a variable and exits the block if a change occurs at the end of any
statement.
```
watch e {
  ...
  e := "error"
  ...
}
```

#### Native Exit Script

Exit the script with a specific exit code:
```
exit exitCode
```

#### Template Strings & Spells

```
a, op, b, eql, c := 1, "+", 2, "=", 3
x := @Fmt("{a} {op} {b} {eql} {c}") // 1 + 2 = 3
x := @fmt.Fmt("{a} {op} {b} {eql} {c}")
```

```
a, b, c := 1.1, 2.22, 3.333
x := @Fmt("{a, .2} + {b, .2} = {c, .2}")
```

#### Native List Accessors

```
x := list[0]
```

#### Function Receivers for Types

```
list := @list.New(1, 2, 3)
list::Len()
```

Or maybe:
```
list := @list.New(1, 2, 3)
list.Len()
```

#### Catch Spell

Executes the code in the block and returns the error if there is a panic.
```
e := @Catch({
  ...
})
```
