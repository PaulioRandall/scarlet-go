
# Package: lexor/symbol

This package was created to separate the concern of managing access to terminal symbols within a stream with the concern of scanning tokens from a script; this package manages access to the terminal symbols. Users of the SymbolStream interface are able to inspect and read off sequences of terminal symbols from a stream, while the implementation keeps track of lines and columns within the streamed text.

The SymbolStream API combines three responsibilities:
- The base functions exposes simple stream functionality such as Len, Empty, IsMatch, CountSymbolsWhile, Peek, and Read.
- The tracking functions, LineIndex and ColIndex, return the position of the stream relative to the text being streamed.
- The remaining functions build upon the base functions to provide slightly higher level capabilities.

Key decisions:
1. The three responsibilities seemed small and simple enough that a single interface combining them would be straight forward to create and maintain as well as being easier for package users to use.
2. The line separator terminals are hardcoded into the package because the program is only expected to work on platforms using those line separators.
3. Any error results in a panic because all errors that can occur are errors made programming to the SymbolStream interface --or errors in the implementation--.
