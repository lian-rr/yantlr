# Y∆NTLR _(Yet ANother Tool for Language Recognition)_


YANTLR is a parser and lexer generator for context free grammars (similar to [ANTLR](https://www.antlr.org)).

Uses an [EBNF](https://en.wikipedia.org/wiki/Extended_Backus–Naur_form) formatted input file and generates (will) the lexer and parser for the grammar.

Input file extension: **.yant**

Use **-help** to see the flags.

## Build

`cd src && go build -o ../yantlr && cd ..`

## Run

`./yantlr -input=examples/pascal.yant`

## File tree

* _src/_: YANTLR code.
* _examples/_: input examples.


## Issues
Well, this doesn't do what it says it does, so, as a friend once said: _open a ticket_.
