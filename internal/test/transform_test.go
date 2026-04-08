package test

import (
	"testing"

	a "acad.learn2earn.ng/git/dositadi/ascii-art-fs/internal/ascii"
)

func TestTransform(t *testing.T) {
	tc := []struct {
		In   []string
		font string
		Want int
	}{
		{In: []string{"Hello", "", "World"}, Want: 3, font: "shadow"},
		{In: []string{"Testing splitter"}, Want: 1, font: "tinkertoy"},
		{In: []string{"Testing", "splitter"}, Want: 2, font: "standard"},
		{In: []string{"\\n\\n"}, Want: 1, font: "standard"},
	}

	a := a.Ascii{Banner: "shadow"}

	for _, c := range tc {
		out, _ := a.Transform(c.In)

		if len(out) != c.Want {
			t.Fatalf("The output lenght(%v) is not equal to the expected %v", len(out), c.Want)
			return
		}
	}
}
