package flag

import (
	"os"
	"strconv"
	"strings"
)

var (
	Args = os.Args[1:]
)

type Type int

const (
	None = iota
	Single
	Double
)

func (t Type) String() string {
	switch t {
	case None:
		return "None"
	case Single:
		return "Single"
	case Double:
		return "Double"
	}

	panic("Unknown value: " + strconv.FormatInt(int64(t), 10))
}

func Walk(f func(val string, t Type) bool) {
	for _, arg := range Args {
		switch {
		case strings.HasPrefix(arg, "--"):
			if !f(arg[2:], Double) {
				return
			}

		case strings.HasPrefix(arg, "-"):
			if arg == "-" {
				if !f(arg, None) {
					return
				}
				continue
			}

			for _, c := range arg[1:] {
				if !f(string(c), Single) {
					return
				}
			}

		default:
			if !f(arg, None) {
				return
			}
		}
	}
}
