package main

type pacmanMode struct{}

var (
	PacmanMode Mode = pacmanMode{}
)

func (m pacmanMode) Main() {
	panic("Not implemented.")
}
