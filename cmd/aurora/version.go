package main

type versionMode struct{}

var VersionMode Mode = &versionMode{}

func (m *versionMode) Main() {
	panic("Not implemented.")
}
