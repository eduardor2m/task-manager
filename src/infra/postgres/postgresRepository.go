package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq"
)

type connectorManager interface {
	getConnection() (*sqlx.DB, error)
	closeConnection(conn *sqlx.DB)
}

var _ connectorManager = (*DatabaseConnectionManager)(nil)

type DatabaseConnectionManager struct{}

func (dcm DatabaseConnectionManager) getConnection() (*sqlx.DB, error) {
	connStr := "postgres://root:root@localhost/task-manager-db?sslmode=disable"
	conn, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	smtp, err := conn.Prepare(
		`CREATE TABLE IF NOT EXISTS task (
			id UUID PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			category VARCHAR(255) NOT NULL,
			status BOOLEAN NOT NULL,
			date TIMESTAMP NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)`,
	)
	if err != nil {
		return nil, err
	}

	_, err = smtp.Exec()

	if err != nil {
		return nil, err
	}

	err = smtp.Close()

	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(10)

	return conn, nil
}

func (dcm DatabaseConnectionManager) closeConnection(conn *sqlx.DB) {
	err := conn.Close()

	if err != nil {
		log.Error(err)
	}
}
