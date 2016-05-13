package flag

import (
	"os"
	"strings"
)

var Args = os.Args[1:]

type Flag interface {
	String() string
}

type Arg string

func (a Arg) String() string {
	return string(a)
}

type Single rune

func (s Single) String() string {
	return "-" + string(s)
}

type Double string

func (d Double) String() string {
	return "--" + string(d)
}

func Walk(f func(Flag) bool) {
	for _, arg := range Args {
		switch {
		case strings.HasPrefix(arg, "--"):
			if !f(Double(arg[2:])) {
				return
			}

		case strings.HasPrefix(arg, "-"):
			if arg == "-" {
				if !f(Arg(arg)) {
					return
				}
				continue
			}

			for _, c := range arg[1:] {
				if !f(Single(c)) {
					return
				}
			}

		default:
			if !f(Arg(arg)) {
				return
			}
		}
	}
}
