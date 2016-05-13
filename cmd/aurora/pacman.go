package main

import (
	"fmt"
	"github.com/DeedleFake/aurora/flag"
	"github.com/DeedleFake/aurora/pacman"
	"os"
	"os/exec"
	"syscall"
)

type pacmanMode struct {
	s bool
	i bool
	p bool

	cmd *exec.Cmd
}

var PacmanMode Mode = &pacmanMode{}

func (m *pacmanMode) modeSudo() {
	m.cmd = Sudo(m.cmd)
}

func (m *pacmanMode) modeNormal() {
}

func (m *pacmanMode) modeSync() {
	if !(m.s || m.i || m.p) {
		m.cmd = Sudo(m.cmd)
	}
}

func (m *pacmanMode) Main() {
	mode := m.modeSudo
	flag.Walk(func(f flag.Flag) bool {
		switch f {
		case flag.Single('Q'):
			mode = m.modeNormal
			return false
		case flag.Single('S'):
			mode = m.modeSync

		case flag.Single('s'), flag.Double("search"):
			m.s = true
		case flag.Single('i'), flag.Double("info"):
			m.i = true
		case flag.Single('p'), flag.Double("print"):
			m.p = true
		}

		return true
	})

	m.cmd = pacman.Pacman(flag.Args...)
	mode()

	err := m.cmd.Run()
	if err != nil {
		switch err := err.(type) {
		case *exec.ExitError:
			os.Exit(err.Sys().(syscall.WaitStatus).ExitStatus())
		default:
			fmt.Printf("aurora: error running pacman: %v", err)
			os.Exit(1)
		}
	}
}
