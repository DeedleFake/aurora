package flag_test

import (
	"."
	"testing"
)

func TestWalk(t *testing.T) {
	flag.Args = []string{"-", "-test", "--test", "test"}
	ex := []struct {
		arg string
		t   flag.Type
	}{
		{"-", flag.None},
		{"t", flag.Single},
		{"e", flag.Single},
		{"s", flag.Single},
		{"t", flag.Single},
		{"test", flag.Double},
		{"test", flag.None},
	}

	i := 0
	flag.Walk(func(arg string, typ flag.Type) bool {
		if (arg != ex[i].arg) || (typ != ex[i].t) {
			t.Errorf("%v: Expected (%q, %v). Got (%q, %v).",
				i,
				ex[i].arg,
				ex[i].t,
				arg,
				typ,
			)
		}

		i++
		return true
	})
}
