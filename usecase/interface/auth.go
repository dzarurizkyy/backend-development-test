package usecase_interface

type AuthUsecase interface {
	Login(username, password string) (string, error)
}
