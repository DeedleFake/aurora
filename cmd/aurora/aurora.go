package main

import (
	"fmt"
	"github.com/DeedleFake/aurora/flag"
	"os"
)

var (
	mode    Mode
	modeMap = map[modeMapKey]Mode{
		{"h", flag.Single}:       HelpMode,
		{"help", flag.Double}:    HelpMode,
		{"V", flag.Single}:       VersionMode,
		{"version", flag.Double}: VersionMode,
		{"D", flag.Single}:       PacmanMode,
		{"R", flag.Single}:       PacmanMode,
		{"F", flag.Single}:       PacmanMode,
		{"Q", flag.Single}:       PacmanMode,
		{"S", flag.Single}:       PacmanMode,
		{"T", flag.Single}:       PacmanMode,
		{"U", flag.Single}:       PacmanMode,
	}
)

type modeMapKey struct {
	arg string
	t   flag.Type
}

func main() {
	flag.Walk(func(arg string, t flag.Type) bool {
		if m, ok := modeMap[modeMapKey{arg, t}]; ok {
			if mode != nil {
				fmt.Println("error: only one operation may be used at a time")
				os.Exit(2)
			}

			mode = m
		}

		return true
	})
	if mode == nil {
		fmt.Println("error: no operation specified (use -h for help)")
		os.Exit(2)
	}

	mode.Main()
}
