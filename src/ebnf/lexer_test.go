package ebnf

import "testing"

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
