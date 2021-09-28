package define

import (
	"testing"
)

const (
	testWord = "sporadic"
)

func TestDictionaryCom(t *testing.T) {
	defs, err := DictionaryCom.Define("sporadic")
	if err != nil {
		t.Fatalf("error getting def: %v", err)
	}

	expected := "(of similar things or occurrences) appearing or happening at irregular intervals in time"

	if defs[0] != expected {
		t.Fatalf("found wrong def for %q: expected %q got %q", DictionaryCom.QueryURL(testWord), expected, defs[0])
	}
}

func TestLexico(t *testing.T) {
	defs, err := Lexico.Define("sporadic")
	if err != nil {
		t.Fatalf("error getting def: %v", err)
	}

	expected := "Occurring at irregular intervals or only in a few places; scattered or isolated."

	if defs[0] != expected {
		t.Fatalf("found wrong def for %q: expected %q got %q", Lexico.QueryURL(testWord), expected, defs[0])
	}
}

func TestCambridge(t *testing.T) {
	defs, err := Cambridge.Define("sporadic")
	if err != nil {
		t.Fatalf("error getting def: %v", err)
	}

	expected := "happening sometimes; not regular or continuous"

	if defs[0] != expected {
		t.Fatalf("found wrong def for %q: expected %q got %q", Cambridge.QueryURL(testWord), expected, defs[0])
	}
}
