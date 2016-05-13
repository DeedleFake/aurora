package main

import (
	"os/exec"
)

var (
	sudoPath string
)

func init() {
	p, err := exec.LookPath("sudo")
	if err != nil {
		panic("sudo not found on local system: " + err.Error())
	}
	sudoPath = p
}

func sudo(cmd *exec.Cmd) *exec.Cmd {
	cmd.Path = sudoPath
	cmd.Args = append([]string{sudoPath}, cmd.Args...)

	return cmd
}
