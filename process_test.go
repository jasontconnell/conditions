package conditions

import (
	"testing"
)

type eqTest struct {
	Name   string
	Result bool
}

func TestEquals(t *testing.T) {
	c := "Name Equals 'Jason'"
	tests := []eqTest{
		{Name: "Jason", Result: true},
		{Name: "Cara", Result: false},
		{Name: "Maura", Result: false},
	}

	condition := Load(c)

	for _, test := range tests {
		if condition.Test(&test) == test.Result {
			t.Log("success")
		} else {
			t.Log("expected true got false", test.Name, c)
			t.Fail()
		}
	}
}

type expTest struct {
	Text  string
	Count int
}

func TestLex(t *testing.T) {
	lexes := []expTest{
		{Text: "Name Equals 'Jason' AND Result Equals true", Count: 7},
		{Text: "(Name NotEquals 'Jason' AND Name NotEquals 'Maura') OR Name Equals 'Steve'", Count: 13},
	}

	for _, lx := range lexes {
		tokens := Lex(lx.Text)

		for _, tk := range tokens {
			t.Log(tk)
		}

		if len(tokens) != lx.Count {
			t.Fail()
		}
	}
}

func TestParse(t *testing.T) {
	lexes := []expTest{
		{Text: "Name Equals 'Jason' AND Result Equals true", Count: 7},
		{Text: "(Name NotEquals 'Jason' AND Name NotEquals 'Maura') OR Name Equals 'Steve'", Count: 13},
	}

	for _, lx := range lexes {
		tokens := Lex(lx.Text)

		exp, err := Parse(tokens)

		if err != nil {
			t.Log(err)
			t.Fail()
		}

		t.Log(exp)
	}
}
