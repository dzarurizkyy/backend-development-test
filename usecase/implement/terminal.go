package usecase_implement

import (
	repository_interface "api-test/repository/interface"
	usecase_interface "api-test/usecase/interface"
)

type terminalUsecase struct {
	repo repository_interface.TerminalRepository
}

func NewTerminalUsecase(r repository_interface.TerminalRepository) usecase_interface.TerminalUsecase {
	return &terminalUsecase{r}
}

func (u *terminalUsecase) CreateTerminal(name string) error {
	return u.repo.CreateTerminal(name)
}
