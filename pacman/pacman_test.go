package pacman_test

import (
	"."
	"testing"
)

func TestPacman(t *testing.T) {
	err := pacman.Pacman("-Qi", "pacman").Run()
	if err != nil {
		t.Error(err)
	}
}
