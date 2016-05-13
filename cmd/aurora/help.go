package main

type helpMode struct{}

var (
	HelpMode Mode = helpMode{}
)

func (m helpMode) Main() {
	panic("Not implemented.")
}
