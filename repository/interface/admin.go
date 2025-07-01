package repository_interface

import "api-test/domain"

type AdminRepository interface {
	GetByUsername(username string) (*domain.Admin, error)
}
