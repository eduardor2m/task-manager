package sqlite

import (
	"database/sql"
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

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm DatabaseConnectionManager) closeConnection(conn *sql.DB) {
	conn.Close()
}
