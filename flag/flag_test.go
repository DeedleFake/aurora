package flag_test

import (
	"."
	"testing"
)

func TestWalk(t *testing.T) {
	flag.Args = []string{"-", "-test", "--test", "test"}
	ex := []struct {
		f flag.Flag
	}{
		{flag.Arg("-")},
		{flag.Single('t')},
		{flag.Single('e')},
		{flag.Single('s')},
		{flag.Single('t')},
		{flag.Double("test")},
		{flag.Arg("test")},
	}

	i := 0
	flag.Walk(func(f flag.Flag) bool {
		if f != ex[i].f {
			t.Errorf("%v: Expected (%T, %[2]q). Got (%T, %[3]q).",
				i,
				ex[i].f,
				f,
			)
		}

		i++
		return true
	})
}
