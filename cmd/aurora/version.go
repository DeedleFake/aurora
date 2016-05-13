package main

import (
	"fmt"
	"github.com/DeedleFake/aurora/pacman"
	"os"
)

var (
	Version = "tip"
)

type versionMode struct{}

var VersionMode Mode = &versionMode{}

func (m *versionMode) Main() {
	fmt.Printf("Aurora %v\n", Version)
	fmt.Println()

	err := pacman.Pacman("-V").Run()
	if err != nil {
		fmt.Printf("aurora: error running pacman: %v\n", err)
		os.Exit(1)
	}
}
