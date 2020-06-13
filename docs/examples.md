
# Exit

```
exit 0
```

```
exit 1
```

```
exit a - 1
```

# Identifiers

```
a
```

# Literals

```
true
```

```
1
```

```
"abc"
```

# Collection Accessors

```
list[1]
```

```
map["abc"]
```

# Negation

```
-1
```

```
-a
```

```
!true
```

```
!a
```

# Existence

```
a?
```

```
f()?
```

# Assignment

```
a := 1
```

```
a := 1 + 1
```

```
a, b := 1, 2
```

```
a, b := f()
```

```
a, b := @s()
```

```
f := F(a -> b) {
	b := a
}
```

```
f := E(a) 1 + 1
```

# Function

```
F() {}
```

```
F(a, b, c) {}
```

```
F(
	a,
	b,
	c,
) {}
```

```
F(a, b -> c, d) {}
```

```
F(a -> a) {}
```

```
F(a, b -> c) {
	c := a + b
}
```

```
F() watch e {
}
```

```
F(a -> b) when a {
	0:       b := "None"
	1:       b := "One"
	[a > 1]: b := "Many"
}
```

```
F(a, b) [a > b] {
}
```

# Expression Function

```
E() 1
```

```
E(a, b, c) a + b - c
```

# Function Call

```
f()
```

```
f(a, b, c)
```

```
f(1 + 1)
```

```
f(
	a,
	b,
	c,
)
```

```
f(a, b, F(c -> d) {
	c := d
})
```

# Spell Call

```
@s()
```

```
@ab.cd()
```

```
@ab.cd.ef()
```

```
@s(a, b, c)
```

```
@s(
	a,
	b,
	c,
)
```

```
@s(a, b, F(c -> d) {
	c := d
})
```

```
@s(1, 2, {
	x := a + b
})
```

# Operation

```
1 + 1
```

```
1 - 1
```

```
1 * 1
```

```
1 / 1
```

```
1 % 1
```

```
1 + 2 - 3 * 4 / 5 % 6
```

```
1 > 1
```

```
1 < 1
```

```
1 >= 1
```

```
1 <= 1
```

```
1 == 1
```

```
1 != 1
```

```
true && false
```

```
true || false
```

```
((1 + 1) * (2 + 2) == 3) && (true && (false || true))
```

# Guards

```
[true] c := 1
```

```
[true] {
	c := 1
}
```

```
[e?] m := "error"
```

```
[a > b] c := 1
```

# Watch

```
watch e {}
```

```
watch a, b, e {}
```

```
watch e {
	a, e := f()
	b := a + 1
}
```

# When

```
when a {}
```

```
when x := 1 {}
```

```
when x := 1 + 1 {}
```

```
when x := true {
	true:  y := 1
	false: y := 2
}
```

```
when x := true {
	true:  {
		y := 1
	}
	false: {
		y := 2
	}
}
```

```
when a {
	0: y := 1
	1: y := 2
	2: y := 3
}
```

```
when x := a {
	[x > 1]: y := 1
	[x < 1]: y := 2
}
```

```
when x := a {
	0:       y := 1
	1:       y := 2
	[x > 1]: y := 3
	2:       y := 4
	[true]:  y := 0
}
```

# Loops

```
loop index := 0 [index < 3] {}
```

```
loop [true] {
	exit loop
}
```

```
loop list -> index, value, hasMore {}
```
