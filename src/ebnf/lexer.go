package ebnf

import (
	"strings"
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
	// Lexer contains the list of tokens.
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
)

//NewLexer initializes and returns a new Lexer
func NewLexer() Lexer {
	return Lexer{
		Tokens: make([]Token, 0),
	}
}

// LoadTokens process all the tokens available in the collection of lines.
func (l *Lexer) LoadTokens(lines []string) {
	ctx := &context{
		open:       false,
		strict:     false,
		line:       0,
		column:     0,
		tokenCount: 0,
	}

	for i, line := range lines {
		tokens := processLine(strings.TrimSuffix(line, " "), i+1, ctx)

		if ctx.open {
			ctx.open = false
			tokens = append(tokens, ctx.buildToken(IDEN))
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
				tokens = append(tokens, ctx.buildToken(IDEN))
			}
			ctx.open = false
			ctx.builder.WriteRune(c)
			ctx.column = i + 1
			tokens = append(tokens, ctx.buildToken(charToTokenType(c)))
			continue
		}

		ctx.builder.WriteRune(c)
	}

	return tokens
}

func isSpecial(c rune) bool {
	switch c {
	case EQUALS, COMMA, SEMICOLON, VERTICAL, DQUOTE, QUOTE, HYPHEN, Q_MARK,
		L_SQUARE_B, R_SQUARE_B, L_CURLY_B, R_CURLY_B, L_PARENTH, R_PARENTH, ASTERISK:
		return true
	default:
		return false
	}

}

func (ctx *context) buildToken(tType TokenType) Token {
	ctx.tokenCount++
	t := Token{
		Id:     ctx.tokenCount,
		Type:   tType,
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
	case EQUALS: // =
		return DEF
	case COMMA: // ,
		return CONCAT
	case SEMICOLON: // ;
		return TERM
	case VERTICAL: // |
		return ALTER
	case L_SQUARE_B, R_SQUARE_B: // [ ]
		return OPT
	case L_CURLY_B, R_CURLY_B: // { }
		return REP
	case L_PARENTH, R_PARENTH: // ( )
		return GROUP
	case QUOTE, DQUOTE: // ' "
		return TERMI
	case Q_MARK: // ?
		return SPEC
	case HYPHEN: // -
		return EXCEP
	default:
		return UNK
	}
}
