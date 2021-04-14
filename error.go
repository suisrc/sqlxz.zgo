package sqlxz

import "strings"

// IsNotFound of sqlx
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return "sql: no rows in result set" == err.Error()
}

func IsDuplicate(err error) bool {
	if err == nil {
		return false
	}
	return strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry ")
}
