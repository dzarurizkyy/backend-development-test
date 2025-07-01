package usecase_implement

import (
	repository_interface "api-test/repository/interface"
	"api-test/utils"
	"errors"

	usecase_interface "api-test/usecase/interface"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	repo repository_interface.AdminRepository
}

func NewAuthUsecase(r repository_interface.AdminRepository) usecase_interface.AuthUsecase {
	return &authUsecase{r}
}

func (a *authUsecase) Login(username, password string) (string, error) {
	admin, err := a.repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if err = bcrypt.CompareHashAndPassword([]uint8(admin.Password), []uint8(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(admin.Id)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
