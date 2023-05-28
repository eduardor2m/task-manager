package validator

import (
	"time"
)

func IsDateValid(date time.Time) bool {
	return date != time.Time{}
}
