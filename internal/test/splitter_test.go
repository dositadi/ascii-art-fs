package test

import (
	"slices"
	"testing"

	a "acad.learn2earn.ng/git/dositadi/ascii-art-fs/internal/ascii"
)

func TestSplitter(t *testing.T) {
	tc := []struct {
		In   string
		Want []string
	}{
		{In: "Hello\\n\\nWorld", Want: []string{"Hello", "", "World"}},
		{In: "Testing splitter", Want: []string{"Testing splitter"}},
		{In: "Testing\\nsplitter", Want: []string{"Testing", "splitter"}},
	}

	for _, c := range tc {
		a := a.Ascii{Input: c.In}
		out := a.Splitter()

		if slices.Compare(out, c.Want) != 0 {
			t.Fatalf("The output %+v is not same as the input.", out)
		}
	}
}
