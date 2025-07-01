package repository_interface

type TerminalRepository interface {
	CreateTerminal(name string) error
}
