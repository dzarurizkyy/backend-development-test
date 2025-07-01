package repository_implement

import (
	repository_interface "api-test/repository/interface"
	"database/sql"
	"time"
)

type terminalRepository struct {
	db *sql.DB
}

func NewTerminalRepository(db *sql.DB) repository_interface.TerminalRepository {
	return &terminalRepository{db}
}

func (t *terminalRepository) CreateTerminal(name string) error {
	query := `INSERT INTO terminal (name, created_at, updated_at, deleted_at, is_deleted) VALUES ($1, $2, $3, $4, $5)`
	_, err := t.db.Exec(query, name, time.Now(), time.Now(), 0, false)
	if err != nil {
		return err
	}

	return nil
}
