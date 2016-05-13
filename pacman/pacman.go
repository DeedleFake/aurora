package pacman

import (
	"os"
	"os/exec"
)

var (
	pacman string
)

func init() {
	p, err := exec.LookPath("pacman")
	if err != nil {
		panic("pacman not found on local system: " + err.Error())
	}
	pacman = p
}

func Run(args ...string) error {
	cmd := &exec.Cmd{
		Path:   pacman,
		Args:   append([]string{pacman}, args...),
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	return cmd.Run()
}
