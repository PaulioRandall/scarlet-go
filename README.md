# scarlet-go

  "Sometimes it's better to light a flamethrower than curse the darkness."
    - 'Men at Arms' by Terry Pratchett

Scarlet is a template for building domain or team specific scripting tools. I started it as a way to learn Go while attempting to create an extremely light weight API testing tool. Another objective is a binary small enough to be included within version control repositories while being easily modifiable at source.

## Syntax

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

# Roadmap

## Next Chapter (Reimplementation)

- Reimplemention using parse trees

## Further Chapters

- Error Handling: Improved error handling and useful error messages
- Spell: Spells are inbuilt functions and will be able to return multiple values
- TinyGo Compatibility: Update so Scarlet can be compiled using TinyGo. This has been done to some degree but I'm waiting on a versioned resolution to https://github.com/tinygo-org/tinygo/issues/890 which is **probably** preventing https://github.com/shopspring/decimal from compiling.
- Common spells
- String spells
- List type and spells
- Map type and spells
- Functions
- Expression functions
