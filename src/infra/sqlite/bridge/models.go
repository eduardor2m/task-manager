// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Completed   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
