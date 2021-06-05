package grammar

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
	t.Logf("\t%v Runes identified as expected", checkMark)
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

	t.Log("Given the need to generate the tokens based on the context.")

	t.Log("\tWhen checking if generated tokens have expected properties")

	for _, r := range tokens {
		r.ctx.builder.WriteString(r.symbol)

		tk := r.ctx.buildToken(r.tType)

		if r.token.Id != tk.Id {
			t.Fatalf("\t\t%v Should have been id: '%d' but actual value is'%d'", ballotX, r.token.Id, tk.Id)
		}

		if r.token.Type != tk.Type {
			t.Fatalf("\t\t%v Should have been token type: '%s' but actual value is'%s'", ballotX, r.token.Type, tk.Type.String())
		}

		if r.token.Position.Line != tk.Position.Line {
			t.Fatalf("\t\t%v Should have been line: '%d' but actual value is'%d'", ballotX, r.token.Position.Line, tk.Position.Line)
		}

		if r.token.Position.Column != tk.Position.Column {
			t.Fatalf("\t\t%v Should have been column: '%d' but actual value is'%d'", ballotX, r.token.Position.Column, tk.Position.Column)
		}
	}

	t.Logf("\t%v Token generated as expected", checkMark)
}

func TestLoadTokensBasicDef(t *testing.T) {
	data := struct {
		line   string
		tokens []Token
	}{
		line: `sentence = "hello world";`,
		tokens: []Token{
			{
				Id:     1,
				Type:   IDEN,
				Symbol: "sentence",
				Position: Position{
					Line: 1, Column: 1,
				},
			},
			{
				Id:     2,
				Type:   DEF,
				Symbol: "=",
				Position: Position{
					Line: 1, Column: 10,
				},
			},
			{
				Id:     3,
				Type:   TERMI,
				Symbol: `"`,
				Position: Position{
					Line: 1, Column: 12,
				},
			},
			{
				Id:     4,
				Type:   IDEN,
				Symbol: `hello world`,
				Position: Position{
					Line: 1, Column: 13,
				},
			},
			{
				Id:     5,
				Type:   TERMI,
				Symbol: `"`,
				Position: Position{
					Line: 1, Column: 24,
				},
			},
			{
				Id:     6,
				Type:   TERM,
				Symbol: `;`,
				Position: Position{
					Line: 1, Column: 25,
				},
			},
		},
	}

	t.Log("Given the need to parse the line for basic definition.")

	lexer := NewLexer()

	lexer.LoadTokens([]string{data.line})

	t.Log("\tWhen checking if generated tokens are the same amount as expected")

	if len(data.tokens) != len(lexer.Tokens) {
		t.Fatalf("\t\t%v Should have been generated '%d' number of tokens but actual value is'%d'",
			ballotX, len(data.tokens), len(lexer.Tokens))
	}

	t.Logf("\t%v Correct number of tokens generated", checkMark)

	t.Log("\tWhen checking if generated tokens have expected properties")

	for i, aToken := range lexer.Tokens {
		eToken := data.tokens[i]

		t.Logf("\t\tWhen checking token %s", aToken.String())

		if aToken.Id != eToken.Id {
			t.Errorf("\t\t\t%v Should have been '%d' as id but actual value is'%d'", ballotX, eToken.Id, aToken.Id)
		}

		if aToken.Type != eToken.Type {
			t.Errorf("\t\t\t%v Should have been '%s' as type but actual value is'%s'", ballotX, eToken.Type.String(), aToken.Type.String())
		}

		if aToken.Symbol != eToken.Symbol {
			t.Errorf("\t\t\t%v Should have been '%s' as symbol but actual value is'%s'", ballotX, eToken.Type.String(), aToken.Type.String())
		}

		if aToken.Position.Line != eToken.Position.Line || aToken.Position.Column != eToken.Position.Column {
			t.Errorf("\t\t\t%v Should have been '%d:%d' as symbol position but actual value is'%d:%d'", ballotX,
				eToken.Position.Line, eToken.Position.Column, aToken.Position.Line, aToken.Position.Column)
		}

	}

	if !t.Failed() {
		t.Logf("\t%v Tokens propertly generated", checkMark)
	}

}
