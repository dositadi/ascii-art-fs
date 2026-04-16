package test

import (
	"testing"

	a "acad.learn2earn.ng/git/dositadi/ascii-art-fs/internal/ascii"
)

func TestReadFile(t *testing.T) {
	a := a.Ascii{Banner: "shadow"}
	tc := []struct {
		In rune
	}{
		{In: 'A'},
		{In: 'a'},
		{In: '2'},
		{In: 'B'},
		{In: 'D'},
		{In: 'f'},
		{In: '0'},
		{In: ' '},
	}

	for _, c := range tc {
		out, err := a.ReadFont(c.In)
		if err != nil {
			t.Fatalf("Error: %s; Detail: %s", err.Error, err.Detail)
		}
		if len(out) != 9 {
			t.Fatalf("Input: %v's lenght does not equal 8. it has lenght %v", string(c.In), len(out))
			return
		}
	}
}
