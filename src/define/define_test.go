package define

import (
	"testing"
)

func TestDictionaryCom(t *testing.T) {
	defs, err := DefineDictionaryCom("sporadic")
	if err != nil {
		t.Fatalf("error getting def: %v", err)
	}

	t.Logf(defs[0])
}
