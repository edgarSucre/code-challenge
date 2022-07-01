package strings_test

import (
	"testing"

	ls "github.com/edgarsucre/challenge/src/strings"
)

func TestReplaceRepetitionsLoop(t *testing.T) {
	cases := []map[string]string{
		{
			"expected": "S2BJK",
			"input":    "SBBJK",
		},
		{
			"expected": "S2BK",
			"input":    "SBBK",
		},
		{
			"expected": "A4BCF2L",
			"input":    "ABBBBCFLL",
		},
	}

	for _, c := range cases {
		got := ls.ReplaceRepetition(c["input"])
		expected := c["expected"]
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestSortCharacters(t *testing.T) {
	cases := []map[string]string{
		{
			"expected": "aabcf",
			"input":    "cafab",
		},
	}

	for _, c := range cases {
		got := ls.SortCharacters(c["input"])
		expected := c["expected"]
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		a string
		b string
		o bool
	}{
		{
			a: "bacdc",
			b: "dcbac",
			o: true,
		},
		{
			a: "bacdc",
			b: "dcbad",
			o: false,
		},
	}

	for _, c := range cases {
		got := ls.AreAnagrams(c.a, c.b)
		expected := c.o
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestMakeAnagram(t *testing.T) {
	cases := []struct {
		a string
		b string
		o int
	}{
		{
			a: "cde",
			b: "dcf",
			o: 2,
		},
		{
			a: "bacdc",
			b: "dcbac",
			o: 0,
		},
		{
			a: "bacdc",
			b: "dcbad",
			o: 2,
		},
		{
			a: "fcrxzwscanmligyxyvym",
			b: "jxwtrhvujlmrpdoqbisbwhmgpmeoke",
			o: 30,
		},
	}

	for _, c := range cases {
		got := ls.MakeAnagram(c.a, c.b)
		expected := c.o
		if got != expected {
			t.Errorf("expected: %v, got: %v", expected, got)
		}
	}
}

func TestIsValid(t *testing.T) {
	o := ls.IsValid("abcdefghhgfedecba")
	if o == "YES" {
		t.Fail()
	}
}

func TestIsPalindrome(t *testing.T) {
	o := ls.IsPalindrome("ava")
	if !o {
		t.Fail()
	}

	o = ls.IsPalindrome("dos")
	if o {
		t.Fail()
	}
}

func TestCountSpecialString(t *testing.T) {
	o := ls.CountSpecialString("aaaa", 4)
	if o != 10 {
		t.Fail()
	}
}

func TestInvertStringRecursive(t *testing.T) {
	o := ls.InvertStringRecursive("prueba")
	if o != "abeurp" {
		t.Fail()
	}

	if o != ls.InvertStringLoop("prueba") {
		t.Fail()
	}
}

func TestConvertDecimalToBinary(t *testing.T) {
	o := ls.ConvertDecimalToBinary(2)
	if o != "10" {
		t.Fail()
	}
}
