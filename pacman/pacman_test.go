package pacman_test

import (
	"."
	"testing"
)

func TestRun(t *testing.T) {
	err := pacman.Run("-Qi", "pacman")
	if err != nil {
		t.Error(err)
	}
}
