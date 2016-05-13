package main

type Mode interface {
	Main()
}

type helpMode struct{}

var (
	HelpMode Mode = helpMode{}
)

func (m helpMode) Main() {
	panic("Not implemented.")
}

type versionMode struct{}

var (
	VersionMode Mode = versionMode{}
)

func (m versionMode) Main() {
	panic("Not implemented.")
}

type pacmanMode struct{}

var (
	PacmanMode Mode = pacmanMode{}
)

func (m pacmanMode) Main() {
	panic("Not implemented.")
}
