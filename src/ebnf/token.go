package ebnf

import (
	"fmt"
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

type (
	// Token is a struct for representing the lexeme
	Token struct {
		Id       int
		Type     TokenType
		Symbol   string
		Position Position
	}

	// TokenType is the type of token
	TokenType int

	// Position contains the token position in the input file
	Position struct {
		Line   int
		Column int
	}
)

// String returns a string description of the token
func (t Token) String() string {
	return fmt.Sprintf("ID: %4d Type: %s Line: %2d[%2d] Symbol: %s", t.Id, t.Type.String(), t.Position.Line, t.Position.Column, t.Symbol)
}

// String returns the text represantation of the TokenType
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
