package sqlite

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (dcm DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./task-manager.db")

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	return db, nil
}

func (dcm DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	conn.Close().Error()
}
