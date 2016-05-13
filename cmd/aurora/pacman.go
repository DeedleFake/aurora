package main

import (
	"github.com/DeedleFake/aurora/flag"
	"github.com/DeedleFake/aurora/pacman"
	"os"
	"os/exec"
	"syscall"
)

type pacmanMode struct{}

var (
	PacmanMode Mode = pacmanMode{}
)

func (m pacmanMode) Main() {
	err := Sudo(pacman.Pacman(flag.Args...)).Run()
	if err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			os.Exit(err.Sys().(syscall.WaitStatus).ExitStatus())
		}
	}
}
