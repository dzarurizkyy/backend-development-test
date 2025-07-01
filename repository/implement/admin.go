package repository_implement

import (
	"api-test/domain"
	"database/sql"

	repository_interface "api-test/repository/interface"
)

type adminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) repository_interface.AdminRepository {
	return &adminRepository{db}
}

func (a *adminRepository) GetByUsername(username string) (*domain.Admin, error) {
	query := `SELECT id, username, fullname, password, created_at, updated_at, deleted_at, is_deleted FROM admin WHERE username=$1 AND is_deleted=$2`
	row := a.db.QueryRow(query, username, false)
	var value domain.Admin

	err := row.Scan(&value.Id, &value.Username, &value.Fullname, &value.Password, &value.CreatedAt, &value.UpdatedAt, &value.DeletedAt, &value.IsDeleted)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
