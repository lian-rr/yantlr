package main

import (
	"fmt"
	"strings"
)

const (
	IDEN TokenType = iota
	DEF
	CONCAT
	TERM
	ALTER
	OPT
	REP
	GROUP
	TERMI
	SPEC
	EXCEP
	UNK
)

const (
	EQUALS     = '='
	COMMA      = ','
	SEMICOLON  = ';'
	VERTICAL   = '|'
	L_SQUARE_B = '['
	R_SQUARE_B = ']'
	L_CURLY_B  = '{'
	R_CURLY_B  = '}'
	L_PARENTH  = '('
	R_PARENTH  = ')'
	QUOTE      = '\''
	DQUOTE     = '"'
	HYPHEN     = '-'
	ASTERISK   = '*'
	Q_MARK     = '?'
	WHITE      = ' '
	LINE_BREAK = '\n'
)

type (
	Lexer struct {
		Tokens []Token
	}

	context struct {
		open       bool
		builder    strings.Builder
		strict     bool
		literal    rune
		column     int
		line       int
		tokenCount int
	}

	Token struct {
		Id       int
		Type     TokenType
		Symbol   string
		Position Position
	}

	TokenType int

	Position struct {
		Line   int
		Column int
	}
)

func NewLexer() Lexer {
	return Lexer{
		Tokens: make([]Token, 0),
	}
}

func (l *Lexer) ProcessLines(lines []string) {
	ctx := &context{
		open:       false,
		strict:     false,
		line:       0,
		column:     0,
		tokenCount: 0,
	}

	for i, line := range lines {
		tokens := processLine(strings.TrimSpace(line), i+1, ctx)

		if ctx.open {
			ctx.open = false
			tokens = append(tokens, ctx.getToken(IDEN))
		}

		l.Tokens = append(l.Tokens, tokens...)
	}
}

func processLine(line string, lineNum int, ctx *context) []Token {
	var tokens = make([]Token, 0)

	rslice := []rune(line)

	ctx.line = lineNum

	for i, c := range rslice {
		if c == WHITE && !ctx.strict {
			continue
		}

		if !ctx.open {
			ctx.open = true
			ctx.column = i + 1
		}

		if isSpecial(c) {
			if c == DQUOTE || c == QUOTE {
				if !ctx.strict {
					ctx.strict = true
					ctx.literal = c
				} else if ctx.literal == c {
					ctx.strict = false
					ctx.literal = '0'
				} else {
					ctx.builder.WriteRune(c)
					continue
				}
			}

			if ctx.open && ctx.builder.Len() != 0 {
				tokens = append(tokens, ctx.getToken(IDEN))
			}
			ctx.open = false
			ctx.builder.WriteRune(c)
			ctx.column = i + 1
			tokens = append(tokens, ctx.getToken(charToTokenType(c)))
			continue
		}

		ctx.builder.WriteRune(c)
	}

	return tokens
}

func (t Token) String() string {
	return fmt.Sprintf("ID: %4d Type: %s Line: %2d[%2d] Symbol: %s", t.Id, t.Type.String(), t.Position.Line, t.Position.Column, t.Symbol)
}

func isSpecial(c rune) bool {
	switch c {
	case EQUALS, COMMA, SEMICOLON, VERTICAL, DQUOTE, QUOTE, HYPHEN, Q_MARK, L_SQUARE_B, R_SQUARE_B, L_CURLY_B, R_CURLY_B, L_PARENTH, R_PARENTH, ASTERISK:
		return true
	default:
		return false
	}

}

func (ctx *context) getToken(ttype TokenType) Token {
	ctx.tokenCount++
	t := Token{
		Id:     ctx.tokenCount,
		Type:   ttype,
		Symbol: strings.TrimSpace(ctx.builder.String()),
		Position: Position{
			Line:   ctx.line,
			Column: ctx.column,
		},
	}

	ctx.builder.Reset()
	return t
}

func charToTokenType(c rune) TokenType {
	switch c {
	case EQUALS:
		return DEF
	case COMMA:
		return CONCAT
	case SEMICOLON:
		return TERM
	case VERTICAL:
		return ALTER
	case L_SQUARE_B, R_SQUARE_B:
		return OPT
	case L_CURLY_B, R_CURLY_B:
		return REP
	case L_PARENTH, R_PARENTH:
		return GROUP
	case QUOTE, DQUOTE:
		return TERMI
	case Q_MARK:
		return SPEC
	case HYPHEN:
		return EXCEP
	default:
		return UNK
	}
}

func (t TokenType) String() string {
	switch t {
	case IDEN:
		return "IDEN"
	case DEF:
		return "DEF"
	case TERM:
		return "TERM"
	case ALTER:
		return "ALTER"
	case OPT:
		return "OPT"
	case REP:
		return "REP"
	case GROUP:
		return "GROUP"
	case TERMI:
		return "TERMI"
	case SPEC:
		return "SPEC"
	case EXCEP:
		return "EXCEP"
	default:
		return "UNK"
	}
}
