package usecase_interface

type TerminalUsecase interface {
	CreateTerminal(name string) error
}