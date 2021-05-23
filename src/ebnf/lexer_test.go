package ebnf

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestIsSpecial(t *testing.T) {
	runes := []struct {
		c rune
		r bool
	}{
		{'=', true},
		{',', true},
		{';', true},
		{'|', true},
		{'[', true},
		{']', true},
		{'{', true},
		{'}', true},
		{'(', true},
		{')', true},
		{'\'', true},
		{'"', true},
		{'-', true},
		{'*', true},
		{'?', true},
		{' ', false},
		{'|', true},
		{'\n', false},
		{'2', false},
		{'p', false},
		{'\t', false},
	}

	t.Log("Given the need to identify the special characters while procesing the line.")

	for _, r := range runes {
		if isSpecial(r.c) != r.r {
			t.Fatalf("\t%v Should have detected '%c' as '%t' for special character", ballotX, r.c, r.r)
		}
	}
	t.Logf("\tRunes identified as expected %v", checkMark)
}

func TestBuildToken(t *testing.T) {

	tokens := []struct {
		ctx    context
		symbol string
		tType  TokenType
		token  Token
	}{
		{
			context{
				tokenCount: 0,
				line:       1,
				column:     1,
			},
			"Hello",
			IDEN,
			Token{
				Id:     1,
				Type:   IDEN,
				Symbol: "Hello",
				Position: Position{
					Line:   1,
					Column: 1,
				},
			},
		},
		{
			context{
				tokenCount: 5,
				line:       20,
				column:     41,
			},
			"=",
			ALTER,
			Token{
				Id:     6,
				Type:   ALTER,
				Symbol: "=",
				Position: Position{
					Line:   20,
					Column: 41,
				},
			},
		},
	}

	t.Log("Given the need to identify generate the tokens with based on the context.")

	for _, r := range tokens {
		r.ctx.builder.WriteString(r.symbol)

		tk := r.ctx.buildToken(r.tType)

		if r.token.Id != tk.Id {
			t.Fatalf("\t%v Should have been id: '%d' but actual '%d'", ballotX, r.token.Id, tk.Id)
		}

		if r.token.Type != tk.Type {
			t.Fatalf("\t%v Should have been token type: '%s' but actual '%s'", ballotX, r.token.Type, tk.Type.String())
		}

		if r.token.Position.Line != tk.Position.Line {
			t.Fatalf("\t%v Should have been line: '%d' but actual '%d'", ballotX, r.token.Position.Line, tk.Position.Line)
		}

		if r.token.Position.Column != tk.Position.Column {
			t.Fatalf("\t%v Should have been column: '%d' but actual '%d'", ballotX, r.token.Position.Column, tk.Position.Column)
		}
	}

	t.Logf("\tToken generated as expected %v", checkMark)
}
