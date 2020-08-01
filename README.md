# scarlet-go

Scarlet is my second attempt at creating an interpreter and is the name of the interpreted language

## 1. Essential Features

#### Arithmetic Operations

```
1 + 2
1 - 2
1 * 2
1 / 2
1 % 2
```

```
1 + 2 * 3 - 4 / 5 % 6
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

## 2. Usability Features

#### Native Assignments

```
x := true
x := 1
x := "abc"
x := y
```

```
x, y, z := 1, 2, 3
```

#### Void Assignee

```
_, y := 1, 2
```

#### Native Variable Deletions

```
x := _
``` 

#### Exit Script

Exit the script with a specific exit code:
```
exit exitCode
```

#### String Spells

Find the length of a string:
```
len := @str.Len(s)
```

Take a slice of a string:
```
x := @str.Slice(s, startIdx, endIdx)
```

Get a specific UTF-8 char from the string, i.e rune:
```
x := @str.Char(s, idx)
```

Get the index of a specific UTF-8 char within the string:
```
x := @str.IndexOf(haystack, needle)
```

Join two strings together:
```
x := @str.Join("abc", "xyz")
```

#### List Spells

Create a new list:
```
list := @list.New(
  1,
  2,
  3,
)
```

Find the length of a list:
```
len := @list.Len(list)
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

Add an item to the end of a list:
```
@list.Add(list, val)
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

#### Map Spells (Requires the list type)

Create a new map:
```
map := @map.New(
  1, "one",
  2, "two",
  3, "three",
)
```

Find the length of a map:
```
len := @map.Len(map)
```

Map a value to a key:
```
@map.Set(map, key, value)
```

Get the value of map entry using its key:
```
x := @map.Get(map, key)
```

Remove an map entry:
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

#### Program Spells

Get the program arguments:
```
x := @Args()
```

Does a variable exist:
```
@Exists("y")
```

#### Functions

```
f := F(a, b -> x, y) {
  ...
}

x, y := f(1, 2)
```

## 3. Nice-to-have Features

#### Variable Existence

Test if a variable exists.
```
x := y?
```

#### When Blocks

A form of match block or switch.

```
when x {
  [x > 0] {
    ...
  }
  [x < 0] {
    ...
  }
  [true] {
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
@foreach(list, {
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

