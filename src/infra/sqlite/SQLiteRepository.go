package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type connectorManager interface {
	getConnection() (*sql.DB, error)
	closeConnection(conn *sql.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (dcm DatabaseConnectionManager) getConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./task-manager.db")
	if err != nil {
		return nil, err
	}
	smtp, err := db.Prepare("CREATE TABLE IF NOT EXISTS tasks(id TEXT PRIMARY KEY, title TEXT, description TEXT, completed INTEGER, created_at TEXT, updated_at TEXT)")

	if err != nil {
		return nil, err
	}

	_, err = smtp.Exec()

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm DatabaseConnectionManager) closeConnection(conn *sql.DB) {
	conn.Close()
}
