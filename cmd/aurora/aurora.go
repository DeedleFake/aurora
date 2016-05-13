package main

import (
	"fmt"
	"github.com/DeedleFake/aurora/flag"
	"os"
)

var (
	mode    Mode
	modeMap = map[flag.Flag]Mode{
		flag.Single('h'):       HelpMode,
		flag.Double("help"):    HelpMode,
		flag.Single('V'):       VersionMode,
		flag.Double("version"): VersionMode,
		flag.Single('D'):       PacmanMode,
		flag.Single('R'):       PacmanMode,
		flag.Single('F'):       PacmanMode,
		flag.Single('Q'):       PacmanMode,
		flag.Single('S'):       PacmanMode,
		flag.Single('T'):       PacmanMode,
		flag.Single('U'):       PacmanMode,
	}
)

type Mode interface {
	Main()
}

func main() {
	flag.Walk(func(f flag.Flag) bool {
		if m, ok := modeMap[f]; ok {
			if m == HelpMode {
				if mode != nil {
					return true
				}
			}

			if mode != nil {
				if mode == HelpMode {
					mode = m
					return true
				}

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
